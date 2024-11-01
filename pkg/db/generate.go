package db

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --target ./ent --feature sql/upsert --feature sql/modifier ./schema
