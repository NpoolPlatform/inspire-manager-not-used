// Code generated by ent, DO NOT EDIT.

package pubsubmessgae

import (
	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the pubsubmessgae type in the database.
	Label = "pubsub_messgae"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldUniqueID holds the string denoting the unique_id field in the database.
	FieldUniqueID = "unique_id"
	// FieldMessageID holds the string denoting the message_id field in the database.
	FieldMessageID = "message_id"
	// FieldSender holds the string denoting the sender field in the database.
	FieldSender = "sender"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// Table holds the table name of the pubsubmessgae in the database.
	Table = "pubsub_messgaes"
)

// Columns holds all SQL columns for pubsubmessgae fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldUniqueID,
	FieldMessageID,
	FieldSender,
	FieldBody,
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
)
