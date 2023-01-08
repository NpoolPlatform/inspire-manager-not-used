package discount

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	val "github.com/NpoolPlatform/message/npool"

	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/discount"

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

var appDate = npool.Discount{
	ID:               uuid.NewString(),
	AppID:            uuid.NewString(),
	Discount:         "7.5",
	Circulation:      "10",
	ReleasedByUserID: uuid.NewString(),
	StartAt:          999,
	DurationDays:     999,
	Message:          uuid.NewString(),
	Name:             uuid.NewString(),
}

var (
	appInfo = npool.DiscountReq{
		ID:               &appDate.ID,
		AppID:            &appDate.AppID,
		Discount:         &appDate.Discount,
		Circulation:      &appDate.Circulation,
		ReleasedByUserID: &appDate.ReleasedByUserID,
		StartAt:          &appDate.StartAt,
		DurationDays:     &appDate.DurationDays,
		Message:          &appDate.Message,
		Name:             &appDate.Name,
	}
)

var info *npool.Discount

func createDiscount(t *testing.T) {
	var err error
	info, err = CreateDiscount(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.CreatedAt = info.CreatedAt
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func createDiscounts(t *testing.T) {
	appDates := []npool.Discount{
		{
			ID:               uuid.NewString(),
			AppID:            uuid.NewString(),
			Discount:         "7.5",
			Circulation:      "10",
			ReleasedByUserID: uuid.NewString(),
			StartAt:          999,
			DurationDays:     999,
			Message:          uuid.NewString(),
			Name:             uuid.NewString(),
		},
		{
			ID:               uuid.NewString(),
			AppID:            uuid.NewString(),
			Discount:         "7.5",
			Circulation:      "10",
			ReleasedByUserID: uuid.NewString(),
			StartAt:          999,
			DurationDays:     999,
			Message:          uuid.NewString(),
			Name:             uuid.NewString(),
		},
	}

	apps := []*npool.DiscountReq{}
	for key := range appDates {
		apps = append(apps, &npool.DiscountReq{
			ID:               &appDates[key].ID,
			AppID:            &appDates[key].AppID,
			Discount:         &appDates[key].Discount,
			Circulation:      &appDates[key].Circulation,
			ReleasedByUserID: &appDates[key].ReleasedByUserID,
			StartAt:          &appDates[key].StartAt,
			DurationDays:     &appDates[key].DurationDays,
			Message:          &appDates[key].Message,
			Name:             &appDates[key].Name,
		})
	}

	infos, err := CreateDiscounts(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateDiscount(t *testing.T) {
	var err error
	info, err = UpdateDiscount(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func getDiscount(t *testing.T) {
	var err error
	info, err = GetDiscount(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getDiscounts(t *testing.T) {
	infos, total, err := GetDiscounts(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &appDate)
	}
}

func getDiscountOnly(t *testing.T) {
	var err error
	info, err = GetDiscountOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func existDiscount(t *testing.T) {
	exist, err := ExistDiscount(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existDiscountConds(t *testing.T) {
	exist, err := ExistDiscountConds(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteDiscount(t *testing.T) {
	info, err := DeleteDiscount(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createDiscount", createDiscount)
	t.Run("createDiscounts", createDiscounts)
	t.Run("getDiscount", getDiscount)
	t.Run("getDiscounts", getDiscounts)
	t.Run("getDiscountOnly", getDiscountOnly)
	t.Run("updateDiscount", updateDiscount)
	t.Run("existDiscount", existDiscount)
	t.Run("existDiscountConds", existDiscountConds)
	t.Run("delete", deleteDiscount)
}
