// Code generated by ent, DO NOT EDIT.

package archivementgeneral

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// GoodID applies equality check predicate on the "good_id" field. It's identical to GoodIDEQ.
func GoodID(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// CoinTypeID applies equality check predicate on the "coin_type_id" field. It's identical to CoinTypeIDEQ.
func CoinTypeID(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// TotalUnits applies equality check predicate on the "total_units" field. It's identical to TotalUnitsEQ.
func TotalUnits(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalUnits), v))
	})
}

// SelfUnits applies equality check predicate on the "self_units" field. It's identical to SelfUnitsEQ.
func SelfUnits(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSelfUnits), v))
	})
}

// TotalAmount applies equality check predicate on the "total_amount" field. It's identical to TotalAmountEQ.
func TotalAmount(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalAmount), v))
	})
}

// SelfAmount applies equality check predicate on the "self_amount" field. It's identical to SelfAmountEQ.
func SelfAmount(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSelfAmount), v))
	})
}

// TotalCommission applies equality check predicate on the "total_commission" field. It's identical to TotalCommissionEQ.
func TotalCommission(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalCommission), v))
	})
}

// SelfCommission applies equality check predicate on the "self_commission" field. It's identical to SelfCommissionEQ.
func SelfCommission(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSelfCommission), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppID)))
	})
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppID)))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserID)))
	})
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserID)))
	})
}

// GoodIDEQ applies the EQ predicate on the "good_id" field.
func GoodIDEQ(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// GoodIDNEQ applies the NEQ predicate on the "good_id" field.
func GoodIDNEQ(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGoodID), v))
	})
}

// GoodIDIn applies the In predicate on the "good_id" field.
func GoodIDIn(vs ...uuid.UUID) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldGoodID), v...))
	})
}

// GoodIDNotIn applies the NotIn predicate on the "good_id" field.
func GoodIDNotIn(vs ...uuid.UUID) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldGoodID), v...))
	})
}

// GoodIDGT applies the GT predicate on the "good_id" field.
func GoodIDGT(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGoodID), v))
	})
}

// GoodIDGTE applies the GTE predicate on the "good_id" field.
func GoodIDGTE(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGoodID), v))
	})
}

// GoodIDLT applies the LT predicate on the "good_id" field.
func GoodIDLT(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGoodID), v))
	})
}

// GoodIDLTE applies the LTE predicate on the "good_id" field.
func GoodIDLTE(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGoodID), v))
	})
}

// GoodIDIsNil applies the IsNil predicate on the "good_id" field.
func GoodIDIsNil() predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldGoodID)))
	})
}

// GoodIDNotNil applies the NotNil predicate on the "good_id" field.
func GoodIDNotNil() predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldGoodID)))
	})
}

// CoinTypeIDEQ applies the EQ predicate on the "coin_type_id" field.
func CoinTypeIDEQ(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDNEQ applies the NEQ predicate on the "coin_type_id" field.
func CoinTypeIDNEQ(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDIn applies the In predicate on the "coin_type_id" field.
func CoinTypeIDIn(vs ...uuid.UUID) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDNotIn applies the NotIn predicate on the "coin_type_id" field.
func CoinTypeIDNotIn(vs ...uuid.UUID) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDGT applies the GT predicate on the "coin_type_id" field.
func CoinTypeIDGT(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDGTE applies the GTE predicate on the "coin_type_id" field.
func CoinTypeIDGTE(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLT applies the LT predicate on the "coin_type_id" field.
func CoinTypeIDLT(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLTE applies the LTE predicate on the "coin_type_id" field.
func CoinTypeIDLTE(v uuid.UUID) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDIsNil applies the IsNil predicate on the "coin_type_id" field.
func CoinTypeIDIsNil() predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCoinTypeID)))
	})
}

// CoinTypeIDNotNil applies the NotNil predicate on the "coin_type_id" field.
func CoinTypeIDNotNil() predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCoinTypeID)))
	})
}

// TotalUnitsEQ applies the EQ predicate on the "total_units" field.
func TotalUnitsEQ(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalUnits), v))
	})
}

// TotalUnitsNEQ applies the NEQ predicate on the "total_units" field.
func TotalUnitsNEQ(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTotalUnits), v))
	})
}

// TotalUnitsIn applies the In predicate on the "total_units" field.
func TotalUnitsIn(vs ...uint32) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTotalUnits), v...))
	})
}

// TotalUnitsNotIn applies the NotIn predicate on the "total_units" field.
func TotalUnitsNotIn(vs ...uint32) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTotalUnits), v...))
	})
}

// TotalUnitsGT applies the GT predicate on the "total_units" field.
func TotalUnitsGT(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTotalUnits), v))
	})
}

// TotalUnitsGTE applies the GTE predicate on the "total_units" field.
func TotalUnitsGTE(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTotalUnits), v))
	})
}

// TotalUnitsLT applies the LT predicate on the "total_units" field.
func TotalUnitsLT(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTotalUnits), v))
	})
}

// TotalUnitsLTE applies the LTE predicate on the "total_units" field.
func TotalUnitsLTE(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTotalUnits), v))
	})
}

// TotalUnitsIsNil applies the IsNil predicate on the "total_units" field.
func TotalUnitsIsNil() predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTotalUnits)))
	})
}

// TotalUnitsNotNil applies the NotNil predicate on the "total_units" field.
func TotalUnitsNotNil() predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTotalUnits)))
	})
}

// SelfUnitsEQ applies the EQ predicate on the "self_units" field.
func SelfUnitsEQ(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSelfUnits), v))
	})
}

// SelfUnitsNEQ applies the NEQ predicate on the "self_units" field.
func SelfUnitsNEQ(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSelfUnits), v))
	})
}

// SelfUnitsIn applies the In predicate on the "self_units" field.
func SelfUnitsIn(vs ...uint32) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSelfUnits), v...))
	})
}

// SelfUnitsNotIn applies the NotIn predicate on the "self_units" field.
func SelfUnitsNotIn(vs ...uint32) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSelfUnits), v...))
	})
}

// SelfUnitsGT applies the GT predicate on the "self_units" field.
func SelfUnitsGT(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSelfUnits), v))
	})
}

// SelfUnitsGTE applies the GTE predicate on the "self_units" field.
func SelfUnitsGTE(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSelfUnits), v))
	})
}

// SelfUnitsLT applies the LT predicate on the "self_units" field.
func SelfUnitsLT(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSelfUnits), v))
	})
}

// SelfUnitsLTE applies the LTE predicate on the "self_units" field.
func SelfUnitsLTE(v uint32) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSelfUnits), v))
	})
}

// SelfUnitsIsNil applies the IsNil predicate on the "self_units" field.
func SelfUnitsIsNil() predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSelfUnits)))
	})
}

// SelfUnitsNotNil applies the NotNil predicate on the "self_units" field.
func SelfUnitsNotNil() predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSelfUnits)))
	})
}

// TotalAmountEQ applies the EQ predicate on the "total_amount" field.
func TotalAmountEQ(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalAmount), v))
	})
}

// TotalAmountNEQ applies the NEQ predicate on the "total_amount" field.
func TotalAmountNEQ(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTotalAmount), v))
	})
}

// TotalAmountIn applies the In predicate on the "total_amount" field.
func TotalAmountIn(vs ...decimal.Decimal) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTotalAmount), v...))
	})
}

// TotalAmountNotIn applies the NotIn predicate on the "total_amount" field.
func TotalAmountNotIn(vs ...decimal.Decimal) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTotalAmount), v...))
	})
}

// TotalAmountGT applies the GT predicate on the "total_amount" field.
func TotalAmountGT(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTotalAmount), v))
	})
}

// TotalAmountGTE applies the GTE predicate on the "total_amount" field.
func TotalAmountGTE(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTotalAmount), v))
	})
}

// TotalAmountLT applies the LT predicate on the "total_amount" field.
func TotalAmountLT(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTotalAmount), v))
	})
}

// TotalAmountLTE applies the LTE predicate on the "total_amount" field.
func TotalAmountLTE(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTotalAmount), v))
	})
}

// TotalAmountIsNil applies the IsNil predicate on the "total_amount" field.
func TotalAmountIsNil() predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTotalAmount)))
	})
}

// TotalAmountNotNil applies the NotNil predicate on the "total_amount" field.
func TotalAmountNotNil() predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTotalAmount)))
	})
}

// SelfAmountEQ applies the EQ predicate on the "self_amount" field.
func SelfAmountEQ(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSelfAmount), v))
	})
}

// SelfAmountNEQ applies the NEQ predicate on the "self_amount" field.
func SelfAmountNEQ(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSelfAmount), v))
	})
}

// SelfAmountIn applies the In predicate on the "self_amount" field.
func SelfAmountIn(vs ...decimal.Decimal) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSelfAmount), v...))
	})
}

// SelfAmountNotIn applies the NotIn predicate on the "self_amount" field.
func SelfAmountNotIn(vs ...decimal.Decimal) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSelfAmount), v...))
	})
}

// SelfAmountGT applies the GT predicate on the "self_amount" field.
func SelfAmountGT(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSelfAmount), v))
	})
}

// SelfAmountGTE applies the GTE predicate on the "self_amount" field.
func SelfAmountGTE(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSelfAmount), v))
	})
}

// SelfAmountLT applies the LT predicate on the "self_amount" field.
func SelfAmountLT(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSelfAmount), v))
	})
}

// SelfAmountLTE applies the LTE predicate on the "self_amount" field.
func SelfAmountLTE(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSelfAmount), v))
	})
}

// SelfAmountIsNil applies the IsNil predicate on the "self_amount" field.
func SelfAmountIsNil() predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSelfAmount)))
	})
}

// SelfAmountNotNil applies the NotNil predicate on the "self_amount" field.
func SelfAmountNotNil() predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSelfAmount)))
	})
}

// TotalCommissionEQ applies the EQ predicate on the "total_commission" field.
func TotalCommissionEQ(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalCommission), v))
	})
}

// TotalCommissionNEQ applies the NEQ predicate on the "total_commission" field.
func TotalCommissionNEQ(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTotalCommission), v))
	})
}

// TotalCommissionIn applies the In predicate on the "total_commission" field.
func TotalCommissionIn(vs ...decimal.Decimal) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTotalCommission), v...))
	})
}

// TotalCommissionNotIn applies the NotIn predicate on the "total_commission" field.
func TotalCommissionNotIn(vs ...decimal.Decimal) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTotalCommission), v...))
	})
}

// TotalCommissionGT applies the GT predicate on the "total_commission" field.
func TotalCommissionGT(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTotalCommission), v))
	})
}

// TotalCommissionGTE applies the GTE predicate on the "total_commission" field.
func TotalCommissionGTE(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTotalCommission), v))
	})
}

// TotalCommissionLT applies the LT predicate on the "total_commission" field.
func TotalCommissionLT(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTotalCommission), v))
	})
}

// TotalCommissionLTE applies the LTE predicate on the "total_commission" field.
func TotalCommissionLTE(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTotalCommission), v))
	})
}

// TotalCommissionIsNil applies the IsNil predicate on the "total_commission" field.
func TotalCommissionIsNil() predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTotalCommission)))
	})
}

// TotalCommissionNotNil applies the NotNil predicate on the "total_commission" field.
func TotalCommissionNotNil() predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTotalCommission)))
	})
}

// SelfCommissionEQ applies the EQ predicate on the "self_commission" field.
func SelfCommissionEQ(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSelfCommission), v))
	})
}

// SelfCommissionNEQ applies the NEQ predicate on the "self_commission" field.
func SelfCommissionNEQ(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSelfCommission), v))
	})
}

// SelfCommissionIn applies the In predicate on the "self_commission" field.
func SelfCommissionIn(vs ...decimal.Decimal) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSelfCommission), v...))
	})
}

// SelfCommissionNotIn applies the NotIn predicate on the "self_commission" field.
func SelfCommissionNotIn(vs ...decimal.Decimal) predicate.ArchivementGeneral {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSelfCommission), v...))
	})
}

// SelfCommissionGT applies the GT predicate on the "self_commission" field.
func SelfCommissionGT(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSelfCommission), v))
	})
}

// SelfCommissionGTE applies the GTE predicate on the "self_commission" field.
func SelfCommissionGTE(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSelfCommission), v))
	})
}

// SelfCommissionLT applies the LT predicate on the "self_commission" field.
func SelfCommissionLT(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSelfCommission), v))
	})
}

// SelfCommissionLTE applies the LTE predicate on the "self_commission" field.
func SelfCommissionLTE(v decimal.Decimal) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSelfCommission), v))
	})
}

// SelfCommissionIsNil applies the IsNil predicate on the "self_commission" field.
func SelfCommissionIsNil() predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSelfCommission)))
	})
}

// SelfCommissionNotNil applies the NotNil predicate on the "self_commission" field.
func SelfCommissionNotNil() predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSelfCommission)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ArchivementGeneral) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ArchivementGeneral) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ArchivementGeneral) predicate.ArchivementGeneral {
	return predicate.ArchivementGeneral(func(s *sql.Selector) {
		p(s.Not())
	})
}
