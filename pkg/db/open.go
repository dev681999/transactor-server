package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"ariga.io/atlas-go-sdk/atlasexec"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.uber.org/zap"

	"transactor-server/pkg/config"
	"transactor-server/pkg/db/ent"
	"transactor-server/pkg/infra/log"
)

// CreateConnStr is a helper method to create a postgres connection url from config
func CreateConnStr(cfg config.DB) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?search_path=%s&sslmode=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.Schema, cfg.SSLMode)
}

// applyMigrations use atlas to apply our migrations
func applyMigrations(ctx context.Context, cfg config.DB) error {
	// Define the execution context, supplying a migration directory
	workdir, err := atlasexec.NewWorkingDir(
		atlasexec.WithMigrations(
			os.DirFS(cfg.MigrationsFolder),
		),
	)
	if err != nil {
		log.L.Error("", zap.Error(err))
		return err
	}
	// atlasexec works on a temporary directory, so we need to close it
	defer workdir.Close()

	// Initialize the client.
	client, err := atlasexec.NewClient(workdir.Path(), "atlas")
	if err != nil {
		log.L.Error("", zap.Error(err))
		return err
	}
	// Run `atlas migrate apply` on a SQLite database under /tmp.
	res, err := client.MigrateApply(ctx, &atlasexec.MigrateApplyParams{
		URL: CreateConnStr(cfg),
	})
	if err != nil {
		log.L.Sugar().Fatalf("failed to apply migrations: %v", err)
		return err
	}

	log.L.Info("Applied migrations", zap.Int("no", len(res.Applied)))

	return nil

}

// openEntClient opens the ent client with pgx dirver wrapped with a open telemetry layer
func openEntClient(cfg config.DB) (*ent.Client, error) {
	connStr := CreateConnStr(cfg)

	// this will wrap our database connect with a open telemetry layer
	// more details here https://github.com/uptrace/opentelemetry-go-extra/tree/main/otelsql
	// this adds a span to passed trace in context and also send certain metrics
	// "go.sql.connections_max_open"
	// "go.sql.connections_open"
	// "go.sql.connections_in_use"
	// "go.sql.connections_idle"
	// "go.sql.connections_wait_count"
	// "go.sql.connections_wait_duration"
	// "go.sql.connections_closed_max_idle"
	// "go.sql.connections_closed_max_idle_time"
	// "go.sql.connections_closed_max_lifetime"
	db, err := otelsql.Open("pgx", connStr,
		otelsql.WithAttributes(semconv.DBSystemPostgreSQL),
		otelsql.WithDBName(cfg.DBName),
	)
	if err != nil {
		log.L.Error("", zap.Error(err))
		return nil, err
	}

	// set a basic setting which is derived from a lot of load testing for medium sized hardware
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)

	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), nil
}

// OpenEntClient applies and pending migration and creates a new ent client
func OpenEntClient(ctx context.Context, cfg config.DB) (*ent.Client, error) {
	err := applyMigrations(ctx, cfg)
	if err != nil {
		return nil, err
	}

	client, err := openEntClient(cfg)
	if err != nil {
		log.L.Error("", zap.Error(err))
		return nil, err
	}

	// this sets ent debug mode which logs all sql queries to console
	if cfg.Debug {
		client = client.Debug()
	}

	return client, nil
}
