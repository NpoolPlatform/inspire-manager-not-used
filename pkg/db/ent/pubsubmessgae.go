// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/pubsubmessgae"
	"github.com/google/uuid"
)

// PubsubMessgae is the model entity for the PubsubMessgae schema.
type PubsubMessgae struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// UniqueID holds the value of the "unique_id" field.
	UniqueID uuid.UUID `json:"unique_id,omitempty"`
	// MessageID holds the value of the "message_id" field.
	MessageID string `json:"message_id,omitempty"`
	// Sender holds the value of the "sender" field.
	Sender string `json:"sender,omitempty"`
	// Body holds the value of the "body" field.
	Body []byte `json:"body,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PubsubMessgae) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case pubsubmessgae.FieldBody:
			values[i] = new([]byte)
		case pubsubmessgae.FieldID, pubsubmessgae.FieldCreatedAt, pubsubmessgae.FieldUpdatedAt, pubsubmessgae.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case pubsubmessgae.FieldMessageID, pubsubmessgae.FieldSender:
			values[i] = new(sql.NullString)
		case pubsubmessgae.FieldUniqueID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PubsubMessgae", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PubsubMessgae fields.
func (pm *PubsubMessgae) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pubsubmessgae.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pm.ID = int(value.Int64)
		case pubsubmessgae.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pm.CreatedAt = uint32(value.Int64)
			}
		case pubsubmessgae.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pm.UpdatedAt = uint32(value.Int64)
			}
		case pubsubmessgae.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pm.DeletedAt = uint32(value.Int64)
			}
		case pubsubmessgae.FieldUniqueID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field unique_id", values[i])
			} else if value != nil {
				pm.UniqueID = *value
			}
		case pubsubmessgae.FieldMessageID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message_id", values[i])
			} else if value.Valid {
				pm.MessageID = value.String
			}
		case pubsubmessgae.FieldSender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sender", values[i])
			} else if value.Valid {
				pm.Sender = value.String
			}
		case pubsubmessgae.FieldBody:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value != nil {
				pm.Body = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this PubsubMessgae.
// Note that you need to call PubsubMessgae.Unwrap() before calling this method if this PubsubMessgae
// was returned from a transaction, and the transaction was committed or rolled back.
func (pm *PubsubMessgae) Update() *PubsubMessgaeUpdateOne {
	return (&PubsubMessgaeClient{config: pm.config}).UpdateOne(pm)
}

// Unwrap unwraps the PubsubMessgae entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pm *PubsubMessgae) Unwrap() *PubsubMessgae {
	_tx, ok := pm.config.driver.(*txDriver)
	if !ok {
		panic("ent: PubsubMessgae is not a transactional entity")
	}
	pm.config.driver = _tx.drv
	return pm
}

// String implements the fmt.Stringer.
func (pm *PubsubMessgae) String() string {
	var builder strings.Builder
	builder.WriteString("PubsubMessgae(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pm.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", pm.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", pm.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", pm.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("unique_id=")
	builder.WriteString(fmt.Sprintf("%v", pm.UniqueID))
	builder.WriteString(", ")
	builder.WriteString("message_id=")
	builder.WriteString(pm.MessageID)
	builder.WriteString(", ")
	builder.WriteString("sender=")
	builder.WriteString(pm.Sender)
	builder.WriteString(", ")
	builder.WriteString("body=")
	builder.WriteString(fmt.Sprintf("%v", pm.Body))
	builder.WriteByte(')')
	return builder.String()
}

// PubsubMessgaes is a parsable slice of PubsubMessgae.
type PubsubMessgaes []*PubsubMessgae

func (pm PubsubMessgaes) config(cfg config) {
	for _i := range pm {
		pm[_i].config = cfg
	}
}