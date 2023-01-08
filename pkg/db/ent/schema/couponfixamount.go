//nolint:dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/mixin"
	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

// CouponFixAmount holds the schema definition for the CouponFixAmount entity.
type CouponFixAmount struct {
	ent.Schema
}

func (CouponFixAmount) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the CouponFixAmount.
func (CouponFixAmount) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("app_id", uuid.UUID{}),
		field.
			Other("denomination", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("circulation", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			UUID("released_by_user_id", uuid.UUID{}),
		field.
			Uint32("start_at").
			Optional().
			Default(0),
		field.
			Uint32("duration_days").
			Optional().
			Default(0),
		field.
			String("message").
			Optional().
			Default(""),
		field.
			String("name").
			Optional().
			Default(""),
		field.
			Uint32("allocated").
			Optional().
			Default(0),
	}
}

// Edges of the CouponFixAmount.
func (CouponFixAmount) Edges() []ent.Edge {
	return nil
}
