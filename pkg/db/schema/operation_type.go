package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// OperationType holds the schema definition for the OperationType entity.
type OperationType struct {
	ent.Schema
}

// Fields of the OperationType.
func (OperationType) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("description"),
		field.Bool("is_debit"),
	}
}

// Edges of the OperationType.
func (OperationType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("transactions", Transaction.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

// Mixin of the OperationType.
func (OperationType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
