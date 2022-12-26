// Code generated by ent, DO NOT EDIT.

package couponfixamount

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the couponfixamount type in the database.
	Label = "coupon_fix_amount"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldDenomination holds the string denoting the denomination field in the database.
	FieldDenomination = "denomination"
	// FieldCirculation holds the string denoting the circulation field in the database.
	FieldCirculation = "circulation"
	// FieldReleaseByUserID holds the string denoting the release_by_user_id field in the database.
	FieldReleaseByUserID = "release_by_user_id"
	// FieldStartAt holds the string denoting the start_at field in the database.
	FieldStartAt = "start_at"
	// FieldDurationDays holds the string denoting the duration_days field in the database.
	FieldDurationDays = "duration_days"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// Table holds the table name of the couponfixamount in the database.
	Table = "coupon_fix_amounts"
)

// Columns holds all SQL columns for couponfixamount fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldAppID,
	FieldDenomination,
	FieldCirculation,
	FieldReleaseByUserID,
	FieldStartAt,
	FieldDurationDays,
	FieldMessage,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/NpoolPlatform/inspire-manager/pkg/db/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultDenomination holds the default value on creation for the "denomination" field.
	DefaultDenomination decimal.Decimal
	// DefaultCirculation holds the default value on creation for the "circulation" field.
	DefaultCirculation decimal.Decimal
	// DefaultStartAt holds the default value on creation for the "start_at" field.
	DefaultStartAt uint32
	// DefaultDurationDays holds the default value on creation for the "duration_days" field.
	DefaultDurationDays uint32
	// DefaultMessage holds the default value on creation for the "message" field.
	DefaultMessage string
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
