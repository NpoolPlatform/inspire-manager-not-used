// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/couponallocated"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/coupondiscount"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/couponfixamount"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/couponspecialoffer"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 4)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   couponallocated.Table,
			Columns: couponallocated.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: couponallocated.FieldID,
			},
		},
		Type: "CouponAllocated",
		Fields: map[string]*sqlgraph.FieldSpec{
			couponallocated.FieldCreatedAt: {Type: field.TypeUint32, Column: couponallocated.FieldCreatedAt},
			couponallocated.FieldUpdatedAt: {Type: field.TypeUint32, Column: couponallocated.FieldUpdatedAt},
			couponallocated.FieldDeletedAt: {Type: field.TypeUint32, Column: couponallocated.FieldDeletedAt},
			couponallocated.FieldAppID:     {Type: field.TypeUUID, Column: couponallocated.FieldAppID},
			couponallocated.FieldUserID:    {Type: field.TypeUUID, Column: couponallocated.FieldUserID},
			couponallocated.FieldType:      {Type: field.TypeString, Column: couponallocated.FieldType},
			couponallocated.FieldCouponID:  {Type: field.TypeUUID, Column: couponallocated.FieldCouponID},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   coupondiscount.Table,
			Columns: coupondiscount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coupondiscount.FieldID,
			},
		},
		Type: "CouponDiscount",
		Fields: map[string]*sqlgraph.FieldSpec{
			coupondiscount.FieldCreatedAt:       {Type: field.TypeUint32, Column: coupondiscount.FieldCreatedAt},
			coupondiscount.FieldUpdatedAt:       {Type: field.TypeUint32, Column: coupondiscount.FieldUpdatedAt},
			coupondiscount.FieldDeletedAt:       {Type: field.TypeUint32, Column: coupondiscount.FieldDeletedAt},
			coupondiscount.FieldAppID:           {Type: field.TypeUUID, Column: coupondiscount.FieldAppID},
			coupondiscount.FieldDiscount:        {Type: field.TypeOther, Column: coupondiscount.FieldDiscount},
			coupondiscount.FieldReleaseByUserID: {Type: field.TypeUUID, Column: coupondiscount.FieldReleaseByUserID},
			coupondiscount.FieldStartAt:         {Type: field.TypeUint32, Column: coupondiscount.FieldStartAt},
			coupondiscount.FieldDurationDays:    {Type: field.TypeUint32, Column: coupondiscount.FieldDurationDays},
			coupondiscount.FieldMessage:         {Type: field.TypeString, Column: coupondiscount.FieldMessage},
			coupondiscount.FieldName:            {Type: field.TypeString, Column: coupondiscount.FieldName},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   couponfixamount.Table,
			Columns: couponfixamount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: couponfixamount.FieldID,
			},
		},
		Type: "CouponFixAmount",
		Fields: map[string]*sqlgraph.FieldSpec{
			couponfixamount.FieldCreatedAt:       {Type: field.TypeUint32, Column: couponfixamount.FieldCreatedAt},
			couponfixamount.FieldUpdatedAt:       {Type: field.TypeUint32, Column: couponfixamount.FieldUpdatedAt},
			couponfixamount.FieldDeletedAt:       {Type: field.TypeUint32, Column: couponfixamount.FieldDeletedAt},
			couponfixamount.FieldAppID:           {Type: field.TypeUUID, Column: couponfixamount.FieldAppID},
			couponfixamount.FieldDenomination:    {Type: field.TypeOther, Column: couponfixamount.FieldDenomination},
			couponfixamount.FieldCirculation:     {Type: field.TypeOther, Column: couponfixamount.FieldCirculation},
			couponfixamount.FieldReleaseByUserID: {Type: field.TypeUUID, Column: couponfixamount.FieldReleaseByUserID},
			couponfixamount.FieldStartAt:         {Type: field.TypeUint32, Column: couponfixamount.FieldStartAt},
			couponfixamount.FieldDurationDays:    {Type: field.TypeUint32, Column: couponfixamount.FieldDurationDays},
			couponfixamount.FieldMessage:         {Type: field.TypeString, Column: couponfixamount.FieldMessage},
			couponfixamount.FieldName:            {Type: field.TypeString, Column: couponfixamount.FieldName},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   couponspecialoffer.Table,
			Columns: couponspecialoffer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: couponspecialoffer.FieldID,
			},
		},
		Type: "CouponSpecialOffer",
		Fields: map[string]*sqlgraph.FieldSpec{
			couponspecialoffer.FieldCreatedAt:       {Type: field.TypeUint32, Column: couponspecialoffer.FieldCreatedAt},
			couponspecialoffer.FieldUpdatedAt:       {Type: field.TypeUint32, Column: couponspecialoffer.FieldUpdatedAt},
			couponspecialoffer.FieldDeletedAt:       {Type: field.TypeUint32, Column: couponspecialoffer.FieldDeletedAt},
			couponspecialoffer.FieldAppID:           {Type: field.TypeUUID, Column: couponspecialoffer.FieldAppID},
			couponspecialoffer.FieldUserID:          {Type: field.TypeUUID, Column: couponspecialoffer.FieldUserID},
			couponspecialoffer.FieldAmount:          {Type: field.TypeOther, Column: couponspecialoffer.FieldAmount},
			couponspecialoffer.FieldReleaseByUserID: {Type: field.TypeUUID, Column: couponspecialoffer.FieldReleaseByUserID},
			couponspecialoffer.FieldStartAt:         {Type: field.TypeUint32, Column: couponspecialoffer.FieldStartAt},
			couponspecialoffer.FieldDurationDays:    {Type: field.TypeUint32, Column: couponspecialoffer.FieldDurationDays},
			couponspecialoffer.FieldMessage:         {Type: field.TypeString, Column: couponspecialoffer.FieldMessage},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (caq *CouponAllocatedQuery) addPredicate(pred func(s *sql.Selector)) {
	caq.predicates = append(caq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CouponAllocatedQuery builder.
func (caq *CouponAllocatedQuery) Filter() *CouponAllocatedFilter {
	return &CouponAllocatedFilter{config: caq.config, predicateAdder: caq}
}

// addPredicate implements the predicateAdder interface.
func (m *CouponAllocatedMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CouponAllocatedMutation builder.
func (m *CouponAllocatedMutation) Filter() *CouponAllocatedFilter {
	return &CouponAllocatedFilter{config: m.config, predicateAdder: m}
}

// CouponAllocatedFilter provides a generic filtering capability at runtime for CouponAllocatedQuery.
type CouponAllocatedFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *CouponAllocatedFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *CouponAllocatedFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(couponallocated.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *CouponAllocatedFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(couponallocated.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *CouponAllocatedFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(couponallocated.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *CouponAllocatedFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(couponallocated.FieldDeletedAt))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *CouponAllocatedFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(couponallocated.FieldAppID))
}

// WhereUserID applies the entql [16]byte predicate on the user_id field.
func (f *CouponAllocatedFilter) WhereUserID(p entql.ValueP) {
	f.Where(p.Field(couponallocated.FieldUserID))
}

// WhereType applies the entql string predicate on the type field.
func (f *CouponAllocatedFilter) WhereType(p entql.StringP) {
	f.Where(p.Field(couponallocated.FieldType))
}

// WhereCouponID applies the entql [16]byte predicate on the coupon_id field.
func (f *CouponAllocatedFilter) WhereCouponID(p entql.ValueP) {
	f.Where(p.Field(couponallocated.FieldCouponID))
}

// addPredicate implements the predicateAdder interface.
func (cdq *CouponDiscountQuery) addPredicate(pred func(s *sql.Selector)) {
	cdq.predicates = append(cdq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CouponDiscountQuery builder.
func (cdq *CouponDiscountQuery) Filter() *CouponDiscountFilter {
	return &CouponDiscountFilter{config: cdq.config, predicateAdder: cdq}
}

// addPredicate implements the predicateAdder interface.
func (m *CouponDiscountMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CouponDiscountMutation builder.
func (m *CouponDiscountMutation) Filter() *CouponDiscountFilter {
	return &CouponDiscountFilter{config: m.config, predicateAdder: m}
}

// CouponDiscountFilter provides a generic filtering capability at runtime for CouponDiscountQuery.
type CouponDiscountFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *CouponDiscountFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *CouponDiscountFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(coupondiscount.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *CouponDiscountFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(coupondiscount.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *CouponDiscountFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(coupondiscount.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *CouponDiscountFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(coupondiscount.FieldDeletedAt))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *CouponDiscountFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(coupondiscount.FieldAppID))
}

// WhereDiscount applies the entql other predicate on the discount field.
func (f *CouponDiscountFilter) WhereDiscount(p entql.OtherP) {
	f.Where(p.Field(coupondiscount.FieldDiscount))
}

// WhereReleaseByUserID applies the entql [16]byte predicate on the release_by_user_id field.
func (f *CouponDiscountFilter) WhereReleaseByUserID(p entql.ValueP) {
	f.Where(p.Field(coupondiscount.FieldReleaseByUserID))
}

// WhereStartAt applies the entql uint32 predicate on the start_at field.
func (f *CouponDiscountFilter) WhereStartAt(p entql.Uint32P) {
	f.Where(p.Field(coupondiscount.FieldStartAt))
}

// WhereDurationDays applies the entql uint32 predicate on the duration_days field.
func (f *CouponDiscountFilter) WhereDurationDays(p entql.Uint32P) {
	f.Where(p.Field(coupondiscount.FieldDurationDays))
}

// WhereMessage applies the entql string predicate on the message field.
func (f *CouponDiscountFilter) WhereMessage(p entql.StringP) {
	f.Where(p.Field(coupondiscount.FieldMessage))
}

// WhereName applies the entql string predicate on the name field.
func (f *CouponDiscountFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(coupondiscount.FieldName))
}

// addPredicate implements the predicateAdder interface.
func (cfaq *CouponFixAmountQuery) addPredicate(pred func(s *sql.Selector)) {
	cfaq.predicates = append(cfaq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CouponFixAmountQuery builder.
func (cfaq *CouponFixAmountQuery) Filter() *CouponFixAmountFilter {
	return &CouponFixAmountFilter{config: cfaq.config, predicateAdder: cfaq}
}

// addPredicate implements the predicateAdder interface.
func (m *CouponFixAmountMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CouponFixAmountMutation builder.
func (m *CouponFixAmountMutation) Filter() *CouponFixAmountFilter {
	return &CouponFixAmountFilter{config: m.config, predicateAdder: m}
}

// CouponFixAmountFilter provides a generic filtering capability at runtime for CouponFixAmountQuery.
type CouponFixAmountFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *CouponFixAmountFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *CouponFixAmountFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(couponfixamount.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *CouponFixAmountFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(couponfixamount.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *CouponFixAmountFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(couponfixamount.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *CouponFixAmountFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(couponfixamount.FieldDeletedAt))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *CouponFixAmountFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(couponfixamount.FieldAppID))
}

// WhereDenomination applies the entql other predicate on the denomination field.
func (f *CouponFixAmountFilter) WhereDenomination(p entql.OtherP) {
	f.Where(p.Field(couponfixamount.FieldDenomination))
}

// WhereCirculation applies the entql other predicate on the circulation field.
func (f *CouponFixAmountFilter) WhereCirculation(p entql.OtherP) {
	f.Where(p.Field(couponfixamount.FieldCirculation))
}

// WhereReleaseByUserID applies the entql [16]byte predicate on the release_by_user_id field.
func (f *CouponFixAmountFilter) WhereReleaseByUserID(p entql.ValueP) {
	f.Where(p.Field(couponfixamount.FieldReleaseByUserID))
}

// WhereStartAt applies the entql uint32 predicate on the start_at field.
func (f *CouponFixAmountFilter) WhereStartAt(p entql.Uint32P) {
	f.Where(p.Field(couponfixamount.FieldStartAt))
}

// WhereDurationDays applies the entql uint32 predicate on the duration_days field.
func (f *CouponFixAmountFilter) WhereDurationDays(p entql.Uint32P) {
	f.Where(p.Field(couponfixamount.FieldDurationDays))
}

// WhereMessage applies the entql string predicate on the message field.
func (f *CouponFixAmountFilter) WhereMessage(p entql.StringP) {
	f.Where(p.Field(couponfixamount.FieldMessage))
}

// WhereName applies the entql string predicate on the name field.
func (f *CouponFixAmountFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(couponfixamount.FieldName))
}

// addPredicate implements the predicateAdder interface.
func (csoq *CouponSpecialOfferQuery) addPredicate(pred func(s *sql.Selector)) {
	csoq.predicates = append(csoq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CouponSpecialOfferQuery builder.
func (csoq *CouponSpecialOfferQuery) Filter() *CouponSpecialOfferFilter {
	return &CouponSpecialOfferFilter{config: csoq.config, predicateAdder: csoq}
}

// addPredicate implements the predicateAdder interface.
func (m *CouponSpecialOfferMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CouponSpecialOfferMutation builder.
func (m *CouponSpecialOfferMutation) Filter() *CouponSpecialOfferFilter {
	return &CouponSpecialOfferFilter{config: m.config, predicateAdder: m}
}

// CouponSpecialOfferFilter provides a generic filtering capability at runtime for CouponSpecialOfferQuery.
type CouponSpecialOfferFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *CouponSpecialOfferFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *CouponSpecialOfferFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(couponspecialoffer.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *CouponSpecialOfferFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(couponspecialoffer.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *CouponSpecialOfferFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(couponspecialoffer.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *CouponSpecialOfferFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(couponspecialoffer.FieldDeletedAt))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *CouponSpecialOfferFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(couponspecialoffer.FieldAppID))
}

// WhereUserID applies the entql [16]byte predicate on the user_id field.
func (f *CouponSpecialOfferFilter) WhereUserID(p entql.ValueP) {
	f.Where(p.Field(couponspecialoffer.FieldUserID))
}

// WhereAmount applies the entql other predicate on the amount field.
func (f *CouponSpecialOfferFilter) WhereAmount(p entql.OtherP) {
	f.Where(p.Field(couponspecialoffer.FieldAmount))
}

// WhereReleaseByUserID applies the entql [16]byte predicate on the release_by_user_id field.
func (f *CouponSpecialOfferFilter) WhereReleaseByUserID(p entql.ValueP) {
	f.Where(p.Field(couponspecialoffer.FieldReleaseByUserID))
}

// WhereStartAt applies the entql uint32 predicate on the start_at field.
func (f *CouponSpecialOfferFilter) WhereStartAt(p entql.Uint32P) {
	f.Where(p.Field(couponspecialoffer.FieldStartAt))
}

// WhereDurationDays applies the entql uint32 predicate on the duration_days field.
func (f *CouponSpecialOfferFilter) WhereDurationDays(p entql.Uint32P) {
	f.Where(p.Field(couponspecialoffer.FieldDurationDays))
}

// WhereMessage applies the entql string predicate on the message field.
func (f *CouponSpecialOfferFilter) WhereMessage(p entql.StringP) {
	f.Where(p.Field(couponspecialoffer.FieldMessage))
}
