//nolint:dupl
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/mixin"
	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

// GoodOrderValuePercent holds the schema definition for the GoodOrderValuePercent entity.
type GoodOrderValuePercent struct {
	ent.Schema
}

func (GoodOrderValuePercent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the GoodOrderValuePercent.
func (GoodOrderValuePercent) Fields() []ent.Field {
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
			Other("percent", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Uint32("start_at").
			Optional().
			Default(uint32(time.Now().Unix())),
		field.
			Uint32("end_at").
			Optional().
			Default(0),
	}
}

// Edges of the GoodOrderValuePercent.
func (GoodOrderValuePercent) Edges() []ent.Edge {
	return nil
}

// Indexes of the GoodOrderValuePercent.
func (GoodOrderValuePercent) Indexes() []ent.Index {
	return nil
}
