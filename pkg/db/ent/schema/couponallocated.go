package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/mixin"

	"github.com/google/uuid"

	"github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
)

// CouponAllocated holds the schema definition for the CouponAllocated entity.
type CouponAllocated struct {
	ent.Schema
}

func (CouponAllocated) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the CouponAllocated.
func (CouponAllocated) Fields() []ent.Field {
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
			String("type").
			Optional().
			Default(allocated.CouponType_DefaultCouponType.String()),
		field.
			UUID("coupon_id", uuid.UUID{}),
	}
}

// Edges of the CouponAllocated.
func (CouponAllocated) Edges() []ent.Edge {
	return nil
}
