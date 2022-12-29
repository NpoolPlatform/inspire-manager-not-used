package detail

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
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/detail"
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

var entity = ent.ArchivementDetail{
	ID:                     uuid.New(),
	AppID:                  uuid.New(),
	UserID:                 uuid.New(),
	GoodID:                 uuid.New(),
	OrderID:                uuid.New(),
	PaymentID:              uuid.New(),
	CoinTypeID:             uuid.New(),
	PaymentCoinTypeID:      uuid.New(),
	PaymentCoinUsdCurrency: decimal.RequireFromString("1.00045000000123012"),
	Amount:                 decimal.RequireFromString("9999999999999999999.999999999999999999"),
	UsdAmount:              decimal.RequireFromString("9999999999999999999.999999999999999999"),
	Units:                  10,
}

var (
	id                     = entity.ID.String()
	appID                  = entity.AppID.String()
	userID                 = entity.UserID.String()
	goodID                 = entity.GoodID.String()
	orderID                = entity.OrderID.String()
	paymentID              = entity.PaymentID.String()
	coinTypeID             = entity.CoinTypeID.String()
	paymentCoinTypeID      = entity.PaymentCoinTypeID.String()
	paymentCoinUSDCurrency = entity.PaymentCoinUsdCurrency.String()
	amount                 = entity.Amount.String()
	usdAmount              = entity.UsdAmount.String()
	units                  = entity.Units

	req = npool.DetailReq{
		ID:                     &id,
		AppID:                  &appID,
		UserID:                 &userID,
		GoodID:                 &goodID,
		OrderID:                &orderID,
		PaymentID:              &paymentID,
		CoinTypeID:             &coinTypeID,
		PaymentCoinTypeID:      &paymentCoinTypeID,
		PaymentCoinUSDCurrency: &paymentCoinUSDCurrency,
		Amount:                 &amount,
		USDAmount:              &usdAmount,
		Units:                  &units,
	}
)

var info *ent.ArchivementDetail

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
	entities := []*ent.ArchivementDetail{
		{
			ID:                     uuid.New(),
			AppID:                  uuid.New(),
			UserID:                 uuid.New(),
			GoodID:                 uuid.New(),
			OrderID:                uuid.New(),
			PaymentID:              uuid.New(),
			CoinTypeID:             uuid.New(),
			PaymentCoinTypeID:      uuid.New(),
			PaymentCoinUsdCurrency: decimal.RequireFromString("1.00045000000123012"),
			Amount:                 decimal.RequireFromString("10.00896"),
			UsdAmount:              decimal.RequireFromString("999123.123142"),
			Units:                  10,
		},
		{
			ID:                     uuid.New(),
			AppID:                  uuid.New(),
			UserID:                 uuid.New(),
			GoodID:                 uuid.New(),
			OrderID:                uuid.New(),
			PaymentID:              uuid.New(),
			CoinTypeID:             uuid.New(),
			PaymentCoinTypeID:      uuid.New(),
			PaymentCoinUsdCurrency: decimal.RequireFromString("1.00045000000123012"),
			Amount:                 decimal.RequireFromString("11.11111"),
			UsdAmount:              decimal.RequireFromString("999123.123142"),
			Units:                  10,
		},
	}

	reqs := []*npool.DetailReq{}
	for _, _entity := range entities {
		_id := _entity.ID.String()
		_appID := _entity.AppID.String()
		_userID := _entity.UserID.String()
		_goodID := _entity.GoodID.String()
		_orderID := _entity.OrderID.String()
		_paymentID := _entity.PaymentID.String()
		_coinTypeID := _entity.CoinTypeID.String()
		_paymentCoinTypeID := entity.PaymentCoinTypeID.String()
		_paymentCoinUSDCurrency := _entity.PaymentCoinUsdCurrency.String()
		_amount := _entity.Amount.String()
		_usdAmount := _entity.UsdAmount.String()
		_units := _entity.Units

		reqs = append(reqs, &npool.DetailReq{
			ID:                     &_id,
			AppID:                  &_appID,
			UserID:                 &_userID,
			GoodID:                 &_goodID,
			OrderID:                &_orderID,
			PaymentID:              &_paymentID,
			CoinTypeID:             &_coinTypeID,
			PaymentCoinTypeID:      &_paymentCoinTypeID,
			PaymentCoinUSDCurrency: &_paymentCoinUSDCurrency,
			Amount:                 &_amount,
			USDAmount:              &_usdAmount,
			Units:                  &_units,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), entity.ID)
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
		if assert.Equal(t, total, 1) {
			assert.Equal(t, infos[0].String(), entity.String())
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

func TestDetail(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	if false {
		t.Run("createBulk", createBulk)
		t.Run("row", row)
		t.Run("rows", rows)
		t.Run("rowOnly", rowOnly)
		t.Run("exist", exist)
		t.Run("existConds", existConds)
		t.Run("count", count)
		t.Run("delete", deleteA)
	}
}
