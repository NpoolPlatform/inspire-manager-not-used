package event

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	alloccoupmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/event"

	testinit "github.com/NpoolPlatform/inspire-manager/pkg/testinit"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = ent.Event{
	ID:        uuid.New(),
	AppID:     uuid.New(),
	EventType: basetypes.UsedFor_Signup.String(),
	Coupons: []npool.Coupon{
		npool.Coupon{ID: uuid.New().String(), CouponType: alloccoupmgrpb.CouponType_FixAmount}, //nolint
		npool.Coupon{ID: uuid.New().String(), CouponType: alloccoupmgrpb.CouponType_Discount},
	},
	Credits:        decimal.RequireFromString("1.798"),
	CreditsPerUsd:  decimal.RequireFromString("1.789"),
	MaxConsecutive: 1,
	InviterLayers:  1,
	GoodID:         uuid.UUID{},
}

var (
	id            = ret.ID.String()
	appID         = ret.AppID.String()
	evType        = basetypes.UsedFor_Signup
	credits       = ret.Credits.String()
	creditsPerUSD = ret.CreditsPerUsd.String()

	req = npool.EventReq{
		ID:            &id,
		AppID:         &appID,
		EventType:     &evType,
		Coupons:       []*npool.Coupon{&ret.Coupons[0], &ret.Coupons[1]},
		Credits:       &credits,
		CreditsPerUSD: &creditsPerUSD,
	}
)

var info *ent.Event

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), ret.String())
	}
}

func createBulk(t *testing.T) {
	entities := []*ent.Event{
		{
			ID:        uuid.New(),
			AppID:     uuid.New(),
			EventType: basetypes.UsedFor_Signin.String(),
			Coupons: []npool.Coupon{
				npool.Coupon{ID: uuid.New().String(), CouponType: alloccoupmgrpb.CouponType_FixAmount}, //nolint
				npool.Coupon{ID: uuid.New().String(), CouponType: alloccoupmgrpb.CouponType_Discount},
			},
			Credits:        decimal.RequireFromString("1.7981"),
			CreditsPerUsd:  decimal.RequireFromString("1.7892"),
			MaxConsecutive: 10,
			InviterLayers:  10,
			GoodID:         uuid.UUID{},
		},
		{
			ID:        uuid.New(),
			AppID:     uuid.New(),
			EventType: basetypes.UsedFor_Signin.String(),
			Coupons: []npool.Coupon{
				npool.Coupon{ID: uuid.New().String(), CouponType: alloccoupmgrpb.CouponType_FixAmount}, //nolint
				npool.Coupon{ID: uuid.New().String(), CouponType: alloccoupmgrpb.CouponType_Discount},
			},
			Credits:        decimal.RequireFromString("1.7983"),
			CreditsPerUsd:  decimal.RequireFromString("1.7894"),
			MaxConsecutive: 11,
			InviterLayers:  11,
			GoodID:         uuid.UUID{},
		},
	}

	reqs := []*npool.EventReq{}

	for _, _event := range entities {
		_id := _event.ID.String()
		_appID := _event.AppID.String()
		_evType := basetypes.UsedFor_Signin
		_credits := _event.Credits.String()
		_creditsPerUSD := _event.CreditsPerUsd.String()

		reqs = append(reqs, &npool.EventReq{
			ID:             &_id,
			AppID:          &_appID,
			EventType:      &_evType,
			Coupons:        []*npool.Coupon{&_event.Coupons[0], &_event.Coupons[1]},
			Credits:        &_credits,
			CreditsPerUSD:  &_creditsPerUSD,
			MaxConsecutive: &_event.MaxConsecutive,
			InviterLayers:  &_event.InviterLayers,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), ret.String())
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&npool.Conds{
			ID: &basetypes.StringVal{Value: id, Op: cruder.EQ},
		}, 0, 0)
	if assert.Nil(t, err) {
		if assert.Equal(t, total, 1) {
			assert.Equal(t, infos[0].String(), ret.String())
		}
	}
}

func rowOnly(t *testing.T) {
	var err error
	info, err = RowOnly(context.Background(),
		&npool.Conds{
			ID: &basetypes.StringVal{Value: id, Op: cruder.EQ},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), ret.String())
	}
}

func count(t *testing.T) {
	count, err := Count(context.Background(),
		&npool.Conds{
			ID: &basetypes.StringVal{Value: id, Op: cruder.EQ},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, uint32(1))
	}
}

func exist(t *testing.T) {
	exist, err := Exist(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existConds(t *testing.T) {
	exist, err := ExistConds(context.Background(),
		&npool.Conds{
			ID: &basetypes.StringVal{Value: id, Op: cruder.EQ},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteA(t *testing.T) {
	info, err := Delete(context.Background(), ret.ID.String())
	if assert.Nil(t, err) {
		ret.DeletedAt = info.DeletedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), ret.String())
	}
}

func TestDetail(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("createBulk", createBulk)
	t.Run("row", row)
	t.Run("rows", rows)
	t.Run("rowOnly", rowOnly)
	t.Run("exist", exist)
	t.Run("existConds", existConds)
	t.Run("count", count)
	t.Run("delete", deleteA)
}
