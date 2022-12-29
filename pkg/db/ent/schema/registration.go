package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/mixin"

	"github.com/google/uuid"
)

// Registration holds the schema definition for the Registration entity.
type Registration struct {
	ent.Schema
}

func (Registration) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Registration.
func (Registration) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("app_id", uuid.UUID{}),
		field.
			UUID("inviter_id", uuid.UUID{}),
		field.
			UUID("invitee_id", uuid.UUID{}),
	}
}

// Edges of the Registration.
func (Registration) Edges() []ent.Edge {
	return nil
}
