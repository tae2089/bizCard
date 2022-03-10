package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// BizCard holds the schema definition for the BizCard entity.
type BizCard struct {
	ent.Schema
}

// Fields of the BizCard.
func (BizCard) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("phone_number"),
		field.String("email"),
		field.Int("age").Positive(),
	}
}

// Edges of the BizCard.
func (BizCard) Edges() []ent.Edge {
	return nil
}
