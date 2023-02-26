package general

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	testinit "github.com/NpoolPlatform/inspire-manager/pkg/testinit"
	valuedef "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/general"
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

var entity = ent.ArchivementGeneral{
	ID:              uuid.New(),
	AppID:           uuid.New(),
	UserID:          uuid.New(),
	GoodID:          uuid.New(),
	CoinTypeID:      uuid.New(),
	TotalAmount:     decimal.NewFromInt(0),
	SelfAmount:      decimal.NewFromInt(0),
	TotalUnits:      0,
	SelfUnits:       0,
	TotalUnitsV1:    decimal.NewFromInt(0),
	SelfUnitsV1:     decimal.NewFromInt(0),
	TotalCommission: decimal.NewFromInt(0),
	SelfCommission:  decimal.NewFromInt(0),
}

var (
	id              = entity.ID.String()
	appID           = entity.AppID.String()
	userID          = entity.UserID.String()
	goodID          = entity.GoodID.String()
	coinTypeID      = entity.CoinTypeID.String()
	totalAmount     = entity.TotalAmount.String()
	selfAmount      = entity.SelfAmount.String()
	totalUnits      = entity.TotalUnitsV1.String()
	selfUnits       = entity.SelfUnitsV1.String()
	totalCommission = entity.TotalCommission.String()
	selfCommission  = entity.SelfCommission.String()

	req = npool.GeneralReq{
		ID:              &id,
		AppID:           &appID,
		UserID:          &userID,
		GoodID:          &goodID,
		CoinTypeID:      &coinTypeID,
		TotalAmount:     &totalAmount,
		SelfAmount:      &selfAmount,
		TotalUnits:      &totalUnits,
		SelfUnits:       &selfUnits,
		TotalCommission: &totalCommission,
		SelfCommission:  &selfCommission,
	}
)

var info *ent.ArchivementGeneral

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &req)
	if assert.Nil(t, err) {
		entity.UpdatedAt = info.UpdatedAt
		entity.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), entity.String())
	}
}

func createBulk(t *testing.T) {
	entities := []*ent.ArchivementGeneral{
		{
			ID:              uuid.New(),
			AppID:           uuid.New(),
			UserID:          uuid.New(),
			GoodID:          uuid.New(),
			CoinTypeID:      uuid.New(),
			TotalAmount:     decimal.NewFromInt(0),
			SelfAmount:      decimal.NewFromInt(0),
			TotalUnits:      1,
			SelfUnits:       1,
			TotalCommission: decimal.NewFromInt(0),
			SelfCommission:  decimal.NewFromInt(0),
			TotalUnitsV1:    decimal.NewFromInt(0),
			SelfUnitsV1:     decimal.NewFromInt(0),
		},
		{
			ID:              uuid.New(),
			AppID:           uuid.New(),
			UserID:          uuid.New(),
			GoodID:          uuid.New(),
			CoinTypeID:      uuid.New(),
			TotalAmount:     decimal.NewFromInt(0),
			SelfAmount:      decimal.NewFromInt(0),
			TotalUnits:      2,
			SelfUnits:       0,
			TotalCommission: decimal.NewFromInt(0),
			SelfCommission:  decimal.NewFromInt(0),
			TotalUnitsV1:    decimal.NewFromInt(0),
			SelfUnitsV1:     decimal.NewFromInt(0),
		},
	}

	reqs := []*npool.GeneralReq{}
	for _, _entity := range entities {
		_id := _entity.ID.String()
		_appID := _entity.AppID.String()
		_userID := _entity.UserID.String()
		_goodID := _entity.GoodID.String()
		_coinTypeID := _entity.CoinTypeID.String()
		_totalAmount := _entity.TotalAmount.String()
		_selfAmount := _entity.SelfAmount.String()
		_totalUnits := _entity.TotalUnitsV1.String()
		_selfUnits := _entity.SelfUnitsV1.String()
		_totalCommission := _entity.TotalCommission.String()
		_selfCommission := _entity.SelfCommission.String()

		reqs = append(reqs, &npool.GeneralReq{
			ID:              &_id,
			AppID:           &_appID,
			UserID:          &_userID,
			GoodID:          &_goodID,
			CoinTypeID:      &_coinTypeID,
			TotalAmount:     &_totalAmount,
			SelfAmount:      &_selfAmount,
			TotalUnits:      &_totalUnits,
			SelfUnits:       &_selfUnits,
			TotalCommission: &_totalCommission,
			SelfCommission:  &_selfCommission,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func add(t *testing.T) {
	totalAmount = "30"
	totalUnits = decimal.NewFromInt(10).String()
	selfUnits = decimal.NewFromInt(10).String()

	req.TotalAmount = &totalAmount
	req.SelfUnits = &selfUnits
	req.TotalUnits = &totalUnits

	entity.TotalAmount, _ = decimal.NewFromString(totalAmount)
	entity.TotalUnitsV1 = decimal.NewFromInt(10)
	entity.SelfUnitsV1 = decimal.NewFromInt(10)

	info, err := AddFields(context.Background(), &req)
	if assert.Nil(t, err) {
		entity.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), entity.String())
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), entity.String())
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
		assert.Equal(t, total, 1)
		assert.Equal(t, infos[0].String(), entity.String())
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
		assert.Equal(t, info.String(), entity.String())
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
	exist, err := Exist(context.Background(), entity.ID)
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
	info, err := Delete(context.Background(), entity.ID)
	if assert.Nil(t, err) {
		entity.DeletedAt = info.DeletedAt
		assert.Equal(t, info.String(), entity.String())
	}
}

func TestGeneral(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("createBulk", createBulk)
	t.Run("add", add)
	t.Run("row", row)
	t.Run("rows", rows)
	t.Run("rowOnly", rowOnly)
	t.Run("exist", exist)
	t.Run("existConds", existConds)
	t.Run("count", count)
	t.Run("delete", deleteA)
}
