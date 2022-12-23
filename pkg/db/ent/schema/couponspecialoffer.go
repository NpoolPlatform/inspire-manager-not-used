package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/mixin"
	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

// CouponSpecialOffer holds the schema definition for the CouponSpecialOffer entity.
type CouponSpecialOffer struct {
	ent.Schema
}

func (CouponSpecialOffer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the CouponSpecialOffer.
func (CouponSpecialOffer) Fields() []ent.Field {
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
			Other("amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			UUID("release_by_user_id", uuid.UUID{}),
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
	}
}

// Edges of the CouponSpecialOffer.
func (CouponSpecialOffer) Edges() []ent.Edge {
	return nil
}
