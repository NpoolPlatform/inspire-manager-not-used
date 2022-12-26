// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/coupondiscount"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CouponDiscount is the model entity for the CouponDiscount schema.
type CouponDiscount struct {
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
	// Discount holds the value of the "discount" field.
	Discount decimal.Decimal `json:"discount,omitempty"`
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
func (*CouponDiscount) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case coupondiscount.FieldDiscount:
			values[i] = new(decimal.Decimal)
		case coupondiscount.FieldCreatedAt, coupondiscount.FieldUpdatedAt, coupondiscount.FieldDeletedAt, coupondiscount.FieldStartAt, coupondiscount.FieldDurationDays:
			values[i] = new(sql.NullInt64)
		case coupondiscount.FieldMessage, coupondiscount.FieldName:
			values[i] = new(sql.NullString)
		case coupondiscount.FieldID, coupondiscount.FieldAppID, coupondiscount.FieldReleaseByUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CouponDiscount", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CouponDiscount fields.
func (cd *CouponDiscount) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case coupondiscount.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				cd.ID = *value
			}
		case coupondiscount.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cd.CreatedAt = uint32(value.Int64)
			}
		case coupondiscount.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cd.UpdatedAt = uint32(value.Int64)
			}
		case coupondiscount.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				cd.DeletedAt = uint32(value.Int64)
			}
		case coupondiscount.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				cd.AppID = *value
			}
		case coupondiscount.FieldDiscount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field discount", values[i])
			} else if value != nil {
				cd.Discount = *value
			}
		case coupondiscount.FieldReleaseByUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field release_by_user_id", values[i])
			} else if value != nil {
				cd.ReleaseByUserID = *value
			}
		case coupondiscount.FieldStartAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start_at", values[i])
			} else if value.Valid {
				cd.StartAt = uint32(value.Int64)
			}
		case coupondiscount.FieldDurationDays:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration_days", values[i])
			} else if value.Valid {
				cd.DurationDays = uint32(value.Int64)
			}
		case coupondiscount.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				cd.Message = value.String
			}
		case coupondiscount.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				cd.Name = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CouponDiscount.
// Note that you need to call CouponDiscount.Unwrap() before calling this method if this CouponDiscount
// was returned from a transaction, and the transaction was committed or rolled back.
func (cd *CouponDiscount) Update() *CouponDiscountUpdateOne {
	return (&CouponDiscountClient{config: cd.config}).UpdateOne(cd)
}

// Unwrap unwraps the CouponDiscount entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cd *CouponDiscount) Unwrap() *CouponDiscount {
	_tx, ok := cd.config.driver.(*txDriver)
	if !ok {
		panic("ent: CouponDiscount is not a transactional entity")
	}
	cd.config.driver = _tx.drv
	return cd
}

// String implements the fmt.Stringer.
func (cd *CouponDiscount) String() string {
	var builder strings.Builder
	builder.WriteString("CouponDiscount(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cd.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", cd.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", cd.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", cd.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", cd.AppID))
	builder.WriteString(", ")
	builder.WriteString("discount=")
	builder.WriteString(fmt.Sprintf("%v", cd.Discount))
	builder.WriteString(", ")
	builder.WriteString("release_by_user_id=")
	builder.WriteString(fmt.Sprintf("%v", cd.ReleaseByUserID))
	builder.WriteString(", ")
	builder.WriteString("start_at=")
	builder.WriteString(fmt.Sprintf("%v", cd.StartAt))
	builder.WriteString(", ")
	builder.WriteString("duration_days=")
	builder.WriteString(fmt.Sprintf("%v", cd.DurationDays))
	builder.WriteString(", ")
	builder.WriteString("message=")
	builder.WriteString(cd.Message)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(cd.Name)
	builder.WriteByte(')')
	return builder.String()
}

// CouponDiscounts is a parsable slice of CouponDiscount.
type CouponDiscounts []*CouponDiscount

func (cd CouponDiscounts) config(cfg config) {
	for _i := range cd {
		cd[_i].config = cfg
	}
}
