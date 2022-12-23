package mixin

import (
	"entgo.io/ent"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/privacy"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/rule"
)

func (TimeMixin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (TimeMixin) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rule.FilterTimeRule(),
		},
	}
}
