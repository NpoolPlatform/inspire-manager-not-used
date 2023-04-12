package goodordervaluepercent

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	testinit "github.com/NpoolPlatform/inspire-manager/pkg/testinit"
	valuedef "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission/goodordervaluepercent"

	"github.com/shopspring/decimal"

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

var ret = ent.GoodOrderValuePercent{
	ID:      uuid.New(),
	AppID:   uuid.New(),
	UserID:  uuid.New(),
	GoodID:  uuid.New(),
	Percent: decimal.RequireFromString("1.1"),
}

var (
	id      = ret.ID.String()
	appID   = ret.AppID.String()
	userID  = ret.UserID.String()
	goodID  = ret.GoodID.String()
	percent = ret.Percent.String()

	req = npool.OrderValuePercentReq{
		ID:      &id,
		AppID:   &appID,
		UserID:  &userID,
		GoodID:  &goodID,
		Percent: &percent,
	}
)

func create(t *testing.T) {
	info, err := Create(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.StartAt = info.StartAt
		assert.Equal(t, info.String(), ret.String())
	}
}

func createBulk(t *testing.T) {
	entities := []*ent.GoodOrderValuePercent{
		{
			ID:      uuid.New(),
			AppID:   uuid.New(),
			UserID:  uuid.New(),
			GoodID:  uuid.New(),
			Percent: decimal.RequireFromString("1.2"),
		},
		{
			ID:      uuid.New(),
			AppID:   uuid.New(),
			UserID:  uuid.New(),
			GoodID:  uuid.New(),
			Percent: decimal.RequireFromString("1.3"),
		},
	}

	reqs := []*npool.OrderValuePercentReq{}
	for _, _ret := range entities {
		_id := _ret.ID.String()
		_appID := _ret.AppID.String()
		_userID := _ret.UserID.String()
		_goodID := _ret.GoodID.String()
		_percent := _ret.Percent.String()

		reqs = append(reqs, &npool.OrderValuePercentReq{
			ID:      &_id,
			AppID:   &_appID,
			UserID:  &_userID,
			GoodID:  &_goodID,
			Percent: &_percent,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func update(t *testing.T) {
	percent := "1.2"

	ret.Percent = decimal.RequireFromString(percent)
	req.Percent = &percent

	info, err := Update(context.Background(), &req)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), ret.String())
	}
}

func row(t *testing.T) {
	info, err := Row(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), ret.String())
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
			assert.Equal(t, infos[0].String(), ret.String())
		}
	}
}

func rowOnly(t *testing.T) {
	info, err := RowOnly(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), ret.String())
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
	exist, err := Exist(context.Background(), ret.ID)
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
	info, err := Delete(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		ret.DeletedAt = info.DeletedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), ret.String())
	}
}

func TestPercent(t *testing.T) {
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
