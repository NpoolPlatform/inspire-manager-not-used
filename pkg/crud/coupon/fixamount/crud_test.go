package allocated

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	valuedef "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/inspire/coupon/fixamount"

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

var coupon = ent.CouponFixAmount{
	ID:              uuid.New(),
	AppID:           uuid.New(),
	Denomination:    decimal.NewFromInt(999),
	Circulation:     decimal.NewFromInt(999),
	ReleaseByUserID: uuid.New(),
	StartAt:         999999,
	DurationDays:    999999,
	Message:         uuid.NewString(),
	Name:            uuid.NewString(),
}

var (
	id              = coupon.ID.String()
	appID           = coupon.AppID.String()
	denomination    = coupon.Denomination.String()
	circulation     = coupon.Circulation.String()
	releaseByUserID = coupon.ReleaseByUserID.String()

	req = npool.FixAmountReq{
		ID:              &id,
		AppID:           &appID,
		Denomination:    &denomination,
		Circulation:     &circulation,
		ReleaseByUserID: &releaseByUserID,
		StartAt:         &coupon.StartAt,
		DurationDays:    &coupon.DurationDays,
		Message:         &coupon.Message,
		Name:            &coupon.Name,
	}
)

var info *ent.CouponFixAmount

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
	entities := []*ent.CouponFixAmount{
		{
			ID:              uuid.New(),
			AppID:           uuid.New(),
			Denomination:    decimal.NewFromInt(999),
			Circulation:     decimal.NewFromInt(999),
			ReleaseByUserID: uuid.New(),
			StartAt:         999999,
			DurationDays:    999999,
			Message:         uuid.NewString(),
			Name:            uuid.NewString(),
		},
		{
			ID:              uuid.New(),
			AppID:           uuid.New(),
			Denomination:    decimal.NewFromInt(999),
			Circulation:     decimal.NewFromInt(999),
			ReleaseByUserID: uuid.New(),
			StartAt:         999999,
			DurationDays:    999999,
			Message:         uuid.NewString(),
			Name:            uuid.NewString(),
		},
	}

	reqs := []*npool.FixAmountReq{}
	for _, _coupon := range entities {
		_id := _coupon.ID.String()
		_appID := _coupon.AppID.String()
		_denomination := _coupon.Denomination.String()
		_circulation := _coupon.Circulation.String()
		_releaseByUserID := _coupon.ReleaseByUserID.String()

		reqs = append(reqs, &npool.FixAmountReq{
			ID:              &_id,
			AppID:           &_appID,
			Denomination:    &_denomination,
			Circulation:     &_circulation,
			ReleaseByUserID: &_releaseByUserID,
			StartAt:         &_coupon.StartAt,
			DurationDays:    &_coupon.DurationDays,
			Message:         &_coupon.Message,
			Name:            &_coupon.Name,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &req)
	if assert.Nil(t, err) {
		coupon.UpdatedAt = info.UpdatedAt
		coupon.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), coupon.String())
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
	t.Run("update", update)
	t.Run("row", row)
	t.Run("rows", rows)
	t.Run("rowOnly", rowOnly)
	t.Run("exist", exist)
	t.Run("existConds", existConds)
	t.Run("count", count)
	t.Run("delete", deleteA)
}
