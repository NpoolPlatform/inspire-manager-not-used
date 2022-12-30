// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArchivementDetailsColumns holds the columns for the "archivement_details" table.
	ArchivementDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "direct_contributor_id", Type: field.TypeUUID, Nullable: true},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "order_id", Type: field.TypeUUID, Nullable: true},
		{Name: "self_order", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "payment_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "payment_coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "payment_coin_usd_currency", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "units", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "usd_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "commission", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// ArchivementDetailsTable holds the schema information for the "archivement_details" table.
	ArchivementDetailsTable = &schema.Table{
		Name:       "archivement_details",
		Columns:    ArchivementDetailsColumns,
		PrimaryKey: []*schema.Column{ArchivementDetailsColumns[0]},
	}
	// ArchivementGeneralsColumns holds the columns for the "archivement_generals" table.
	ArchivementGeneralsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "total_units", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "self_units", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "total_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "self_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "total_commission", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "self_commission", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// ArchivementGeneralsTable holds the schema information for the "archivement_generals" table.
	ArchivementGeneralsTable = &schema.Table{
		Name:       "archivement_generals",
		Columns:    ArchivementGeneralsColumns,
		PrimaryKey: []*schema.Column{ArchivementGeneralsColumns[0]},
	}
	// CouponAllocatedsColumns holds the columns for the "coupon_allocateds" table.
	CouponAllocatedsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "type", Type: field.TypeString, Nullable: true, Default: "DefaultCouponType"},
		{Name: "coupon_id", Type: field.TypeUUID},
		{Name: "value", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "used", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "used_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "used_by_order_id", Type: field.TypeUUID, Nullable: true},
	}
	// CouponAllocatedsTable holds the schema information for the "coupon_allocateds" table.
	CouponAllocatedsTable = &schema.Table{
		Name:       "coupon_allocateds",
		Columns:    CouponAllocatedsColumns,
		PrimaryKey: []*schema.Column{CouponAllocatedsColumns[0]},
	}
	// CouponDiscountsColumns holds the columns for the "coupon_discounts" table.
	CouponDiscountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "discount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "circulation", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "released_by_user_id", Type: field.TypeUUID},
		{Name: "start_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "duration_days", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "message", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "allocated", Type: field.TypeUint32, Nullable: true, Default: 0},
	}
	// CouponDiscountsTable holds the schema information for the "coupon_discounts" table.
	CouponDiscountsTable = &schema.Table{
		Name:       "coupon_discounts",
		Columns:    CouponDiscountsColumns,
		PrimaryKey: []*schema.Column{CouponDiscountsColumns[0]},
	}
	// CouponFixAmountsColumns holds the columns for the "coupon_fix_amounts" table.
	CouponFixAmountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "denomination", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "circulation", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "released_by_user_id", Type: field.TypeUUID},
		{Name: "start_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "duration_days", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "message", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "allocated", Type: field.TypeUint32, Nullable: true, Default: 0},
	}
	// CouponFixAmountsTable holds the schema information for the "coupon_fix_amounts" table.
	CouponFixAmountsTable = &schema.Table{
		Name:       "coupon_fix_amounts",
		Columns:    CouponFixAmountsColumns,
		PrimaryKey: []*schema.Column{CouponFixAmountsColumns[0]},
	}
	// CouponSpecialOffersColumns holds the columns for the "coupon_special_offers" table.
	CouponSpecialOffersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "released_by_user_id", Type: field.TypeUUID},
		{Name: "start_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "duration_days", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "message", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// CouponSpecialOffersTable holds the schema information for the "coupon_special_offers" table.
	CouponSpecialOffersTable = &schema.Table{
		Name:       "coupon_special_offers",
		Columns:    CouponSpecialOffersColumns,
		PrimaryKey: []*schema.Column{CouponSpecialOffersColumns[0]},
	}
	// GoodOrderPercentsColumns holds the columns for the "good_order_percents" table.
	GoodOrderPercentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "percent", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "start_at", Type: field.TypeUint32, Nullable: true, Default: 1672389889},
		{Name: "end_at", Type: field.TypeUint32, Nullable: true, Default: 0},
	}
	// GoodOrderPercentsTable holds the schema information for the "good_order_percents" table.
	GoodOrderPercentsTable = &schema.Table{
		Name:       "good_order_percents",
		Columns:    GoodOrderPercentsColumns,
		PrimaryKey: []*schema.Column{GoodOrderPercentsColumns[0]},
	}
	// InvitationCodesColumns holds the columns for the "invitation_codes" table.
	InvitationCodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "invitation_code", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "confirmed", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// InvitationCodesTable holds the schema information for the "invitation_codes" table.
	InvitationCodesTable = &schema.Table{
		Name:       "invitation_codes",
		Columns:    InvitationCodesColumns,
		PrimaryKey: []*schema.Column{InvitationCodesColumns[0]},
	}
	// RegistrationsColumns holds the columns for the "registrations" table.
	RegistrationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "inviter_id", Type: field.TypeUUID},
		{Name: "invitee_id", Type: field.TypeUUID},
	}
	// RegistrationsTable holds the schema information for the "registrations" table.
	RegistrationsTable = &schema.Table{
		Name:       "registrations",
		Columns:    RegistrationsColumns,
		PrimaryKey: []*schema.Column{RegistrationsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArchivementDetailsTable,
		ArchivementGeneralsTable,
		CouponAllocatedsTable,
		CouponDiscountsTable,
		CouponFixAmountsTable,
		CouponSpecialOffersTable,
		GoodOrderPercentsTable,
		InvitationCodesTable,
		RegistrationsTable,
	}
)

func init() {
}
