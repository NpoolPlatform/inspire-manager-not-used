package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/mixin"
	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

// OrderPercent holds the schema definition for the OrderPercent entity.
type OrderPercent struct {
	ent.Schema
}

func (OrderPercent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the OrderPercent.
func (OrderPercent) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			UUID("good_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			String("title").
			Optional().
			Default(""),
		field.
			Other("percent", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Uint32("start_at"),
		field.
			Uint32("end_at"),
		field.
			String("badge_large"),
		field.
			String("badge_small"),
	}
}

// Edges of the OrderPercent.
func (OrderPercent) Edges() []ent.Edge {
	return nil
}

// Indexes of the OrderPercent.
func (OrderPercent) Indexes() []ent.Index {
	return nil
}
