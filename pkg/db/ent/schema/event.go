package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/mixin"
	"github.com/shopspring/decimal"

	"github.com/google/uuid"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

func (Event) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("app_id", uuid.UUID{}),
		field.
			String("event_type").
			Optional().
			Default(basetypes.UsedFor_DefaultUsedFor.String()),
		field.
			JSON("coupon_ids", []uuid.UUID{}).
			Optional().
			Default([]uuid.UUID{}),
		field.
			Other("credits", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("credits_per_usd", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Uint32("max_consecutive").
			Optional().
			Default(1),
		field.
			UUID("good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return nil
}
