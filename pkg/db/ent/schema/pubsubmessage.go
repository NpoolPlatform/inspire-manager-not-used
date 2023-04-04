package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/mixin"
	"github.com/google/uuid"
)

// PubsubMessage holds the schema definition for the PubsubMessage entity.
type PubsubMessage struct {
	ent.Schema
}

func (PubsubMessage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the PubsubMessage.
func (PubsubMessage) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Unique(),
		field.
			String("message_id"),
		field.
			String("state"),
		field.
			UUID("response_to_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
	}
}

// Edges of the PubsubMessage.
func (PubsubMessage) Edges() []ent.Edge {
	return nil
}

func (PubsubMessage) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("state", "response_to_id"),
	}
}
