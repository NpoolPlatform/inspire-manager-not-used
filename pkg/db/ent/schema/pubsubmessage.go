package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/mixin"
	"github.com/google/uuid"
)

// PubsubMessgae holds the schema definition for the PubsubMessgae entity.
type PubsubMessgae struct {
	ent.Schema
}

func (PubsubMessgae) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the PubsubMessgae.
func (PubsubMessgae) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Unique(),
		field.
			String("message_id"),
		field.
			String("sender"),
		field.
			Bytes("body"),
		field.
			String("state"),
		field.
			UUID("response_id", uuid.UUID{}).
			Unique(),
		field.
			String("error_message").
			Optional().
			Default(""),
	}
}

// Edges of the PubsubMessgae.
func (PubsubMessgae) Edges() []ent.Edge {
	return nil
}

func (PubsubMessgae) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("message_id", "response_id"),
	}
}
