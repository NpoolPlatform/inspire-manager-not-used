package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/mixin"

	"github.com/google/uuid"
)

// InvitationCode holds the schema definition for the InvitationCode entity.
type InvitationCode struct {
	ent.Schema
}

func (InvitationCode) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the InvitationCode.
func (InvitationCode) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("app_id", uuid.UUID{}),
		field.
			UUID("user_id", uuid.UUID{}),
		field.
			String("invitation_code").
			Optional().
			Default(""),
		field.
			Bool("confirmed").
			Optional().
			Default(false),
		field.
			Bool("disabled").
			Optional().
			Default(false),
	}
}

// Edges of the InvitationCode.
func (InvitationCode) Edges() []ent.Edge {
	return nil
}
