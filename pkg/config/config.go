package config

type Server struct {
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	APIKey          string `yaml:"api_key"`
	Debug           bool   `yaml:"debug"`
	EnableTelemetry bool   `yaml:"enable_telemetry"`
	OTELEndpoint    string `yaml:"otel_endpoint"`
}

type DB struct {
	Host             string `yaml:"host"`
	Port             int    `yaml:"port"`
	User             string `yaml:"user"`
	DBName           string `yaml:"db_name"`
	Password         string `yaml:"password"`
	Schema           string `yaml:"schema"`
	SSLMode          string `yaml:"ssl_mode"`
	Debug            bool   `yaml:"debug"`
	MigrationsFolder string `yaml:"migrations_folder"`
}

type Config struct {
	Server Server `yaml:"server"`
	DB     DB     `yaml:"db"`
}

const AppName string = "transactor-server"
