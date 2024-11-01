package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Fields of the Transactions.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		// this is a generated uuid
		field.Int("id"),
		field.Int("account_id").Immutable(),
		field.Float("amount").Immutable(),
		field.Int("operation_type_id").Immutable(),
		field.Time("timestamp").Immutable(),
	}
}

// Indexes of the Transactions.
func (Transaction) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("account_id"),
		index.Fields("account_id", "operation_type_id"),
		index.Fields("account_id", "timestamp"),
		index.Fields("account_id", "operation_type_id", "timestamp"),
	}
}

// Edges of the Transactions.
func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("account", Account.Type).
			Field("account_id").
			Ref("transactions").
			Required().
			Immutable().
			Unique(),
		edge.
			From("operation_type", OperationType.Type).
			Field("operation_type_id").
			Ref("transactions").
			Required().
			Immutable().
			Unique(),
	}
}

// Mixin of the Transactions.
func (Transaction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{}, // this is just for audit purposes
	}
}
