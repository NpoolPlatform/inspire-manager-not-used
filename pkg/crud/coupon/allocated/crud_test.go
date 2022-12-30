package allocated

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	valuedef "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"

	testinit "github.com/NpoolPlatform/inspire-manager/pkg/testinit"
	"github.com/google/uuid"

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

var coupon = ent.CouponAllocated{
	ID:         uuid.New(),
	AppID:      uuid.New(),
	UserID:     uuid.New(),
	CouponType: npool.CouponType_Discount.String(),
	CouponID:   uuid.New(),
}

var (
	id       = coupon.ID.String()
	appID    = coupon.AppID.String()
	userID   = coupon.UserID.String()
	cType    = npool.CouponType_Discount
	couponID = coupon.CouponID.String()
	req      = npool.AllocatedReq{
		ID:         &id,
		AppID:      &appID,
		UserID:     &userID,
		CouponType: &cType,
		CouponID:   &couponID,
	}
)

var info *ent.CouponAllocated

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &req)
	if assert.Nil(t, err) {
		coupon.UpdatedAt = info.UpdatedAt
		coupon.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), coupon.String())
	}
}

func createBulk(t *testing.T) {
	entities := []*ent.CouponAllocated{
		{
			ID:         uuid.New(),
			AppID:      uuid.New(),
			UserID:     uuid.New(),
			CouponType: npool.CouponType_Discount.String(),
			CouponID:   uuid.New(),
		},
		{
			ID:         uuid.New(),
			AppID:      uuid.New(),
			UserID:     uuid.New(),
			CouponType: npool.CouponType_Discount.String(),
			CouponID:   uuid.New(),
		},
	}

	reqs := []*npool.AllocatedReq{}
	for _, _coupon := range entities {
		_id := _coupon.ID.String()
		_appID := _coupon.AppID.String()
		_userID := coupon.UserID.String()
		_cType := npool.CouponType_Discount
		_couponID := coupon.CouponID.String()
		reqs = append(reqs, &npool.AllocatedReq{
			ID:         &_id,
			AppID:      &_appID,
			UserID:     &_userID,
			CouponType: &_cType,
			CouponID:   &_couponID,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), coupon.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), coupon.String())
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		if assert.Equal(t, total, 1) {
			assert.Equal(t, infos[0].String(), coupon.String())
		}
	}
}

func rowOnly(t *testing.T) {
	var err error
	info, err = RowOnly(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), coupon.String())
	}
}

func count(t *testing.T) {
	count, err := Count(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, uint32(1))
	}
}

func exist(t *testing.T) {
	exist, err := Exist(context.Background(), coupon.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existConds(t *testing.T) {
	exist, err := ExistConds(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteA(t *testing.T) {
	info, err := Delete(context.Background(), coupon.ID.String())
	if assert.Nil(t, err) {
		coupon.DeletedAt = info.DeletedAt
		coupon.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), coupon.String())
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
