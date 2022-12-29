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

// GoodOrderPercent holds the schema definition for the GoodOrderPercent entity.
type GoodOrderPercent struct {
	ent.Schema
}

func (GoodOrderPercent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the GoodOrderPercent.
func (GoodOrderPercent) Fields() []ent.Field {
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

// Edges of the GoodOrderPercent.
func (GoodOrderPercent) Edges() []ent.Edge {
	return nil
}

// Indexes of the GoodOrderPercent.
func (GoodOrderPercent) Indexes() []ent.Index {
	return nil
}
