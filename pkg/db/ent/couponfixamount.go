// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/couponfixamount"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CouponFixAmount is the model entity for the CouponFixAmount schema.
type CouponFixAmount struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// Denomination holds the value of the "denomination" field.
	Denomination decimal.Decimal `json:"denomination,omitempty"`
	// Circulation holds the value of the "circulation" field.
	Circulation decimal.Decimal `json:"circulation,omitempty"`
	// ReleaseByUserID holds the value of the "release_by_user_id" field.
	ReleaseByUserID uuid.UUID `json:"release_by_user_id,omitempty"`
	// StartAt holds the value of the "start_at" field.
	StartAt uint32 `json:"start_at,omitempty"`
	// DurationDays holds the value of the "duration_days" field.
	DurationDays uint32 `json:"duration_days,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CouponFixAmount) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case couponfixamount.FieldDenomination, couponfixamount.FieldCirculation:
			values[i] = new(decimal.Decimal)
		case couponfixamount.FieldCreatedAt, couponfixamount.FieldUpdatedAt, couponfixamount.FieldDeletedAt, couponfixamount.FieldStartAt, couponfixamount.FieldDurationDays:
			values[i] = new(sql.NullInt64)
		case couponfixamount.FieldMessage, couponfixamount.FieldName:
			values[i] = new(sql.NullString)
		case couponfixamount.FieldID, couponfixamount.FieldAppID, couponfixamount.FieldReleaseByUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CouponFixAmount", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CouponFixAmount fields.
func (cfa *CouponFixAmount) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case couponfixamount.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				cfa.ID = *value
			}
		case couponfixamount.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cfa.CreatedAt = uint32(value.Int64)
			}
		case couponfixamount.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cfa.UpdatedAt = uint32(value.Int64)
			}
		case couponfixamount.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				cfa.DeletedAt = uint32(value.Int64)
			}
		case couponfixamount.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				cfa.AppID = *value
			}
		case couponfixamount.FieldDenomination:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field denomination", values[i])
			} else if value != nil {
				cfa.Denomination = *value
			}
		case couponfixamount.FieldCirculation:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field circulation", values[i])
			} else if value != nil {
				cfa.Circulation = *value
			}
		case couponfixamount.FieldReleaseByUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field release_by_user_id", values[i])
			} else if value != nil {
				cfa.ReleaseByUserID = *value
			}
		case couponfixamount.FieldStartAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start_at", values[i])
			} else if value.Valid {
				cfa.StartAt = uint32(value.Int64)
			}
		case couponfixamount.FieldDurationDays:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration_days", values[i])
			} else if value.Valid {
				cfa.DurationDays = uint32(value.Int64)
			}
		case couponfixamount.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				cfa.Message = value.String
			}
		case couponfixamount.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				cfa.Name = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CouponFixAmount.
// Note that you need to call CouponFixAmount.Unwrap() before calling this method if this CouponFixAmount
// was returned from a transaction, and the transaction was committed or rolled back.
func (cfa *CouponFixAmount) Update() *CouponFixAmountUpdateOne {
	return (&CouponFixAmountClient{config: cfa.config}).UpdateOne(cfa)
}

// Unwrap unwraps the CouponFixAmount entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cfa *CouponFixAmount) Unwrap() *CouponFixAmount {
	_tx, ok := cfa.config.driver.(*txDriver)
	if !ok {
		panic("ent: CouponFixAmount is not a transactional entity")
	}
	cfa.config.driver = _tx.drv
	return cfa
}

// String implements the fmt.Stringer.
func (cfa *CouponFixAmount) String() string {
	var builder strings.Builder
	builder.WriteString("CouponFixAmount(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cfa.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", cfa.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", cfa.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", cfa.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", cfa.AppID))
	builder.WriteString(", ")
	builder.WriteString("denomination=")
	builder.WriteString(fmt.Sprintf("%v", cfa.Denomination))
	builder.WriteString(", ")
	builder.WriteString("circulation=")
	builder.WriteString(fmt.Sprintf("%v", cfa.Circulation))
	builder.WriteString(", ")
	builder.WriteString("release_by_user_id=")
	builder.WriteString(fmt.Sprintf("%v", cfa.ReleaseByUserID))
	builder.WriteString(", ")
	builder.WriteString("start_at=")
	builder.WriteString(fmt.Sprintf("%v", cfa.StartAt))
	builder.WriteString(", ")
	builder.WriteString("duration_days=")
	builder.WriteString(fmt.Sprintf("%v", cfa.DurationDays))
	builder.WriteString(", ")
	builder.WriteString("message=")
	builder.WriteString(cfa.Message)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(cfa.Name)
	builder.WriteByte(')')
	return builder.String()
}

// CouponFixAmounts is a parsable slice of CouponFixAmount.
type CouponFixAmounts []*CouponFixAmount

func (cfa CouponFixAmounts) config(cfg config) {
	for _i := range cfa {
		cfa[_i].config = cfg
	}
}
