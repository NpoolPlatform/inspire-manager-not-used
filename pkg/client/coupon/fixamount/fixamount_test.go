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

	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/fixamount"

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

var appDate = npool.FixAmount{
	ID:              uuid.NewString(),
	AppID:           uuid.NewString(),
	Denomination:    "999999",
	Circulation:     "999999",
	ReleaseByUserID: uuid.NewString(),
	StartAt:         999,
	DurationDays:    999,
	Message:         uuid.NewString(),
	Name:            uuid.NewString(),
}

var (
	appInfo = npool.FixAmountReq{
		ID:              &appDate.ID,
		AppID:           &appDate.AppID,
		Denomination:    &appDate.Denomination,
		Circulation:     &appDate.Circulation,
		ReleaseByUserID: &appDate.ReleaseByUserID,
		StartAt:         &appDate.StartAt,
		DurationDays:    &appDate.DurationDays,
		Message:         &appDate.Message,
		Name:            &appDate.Name,
	}
)

var info *npool.FixAmount

func createFixAmount(t *testing.T) {
	var err error
	info, err = CreateFixAmount(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.CreatedAt = info.CreatedAt
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func createFixAmounts(t *testing.T) {
	appDates := []npool.FixAmount{
		{
			ID:              uuid.NewString(),
			AppID:           uuid.NewString(),
			Denomination:    "999999",
			Circulation:     "999999",
			ReleaseByUserID: uuid.NewString(),
			StartAt:         999,
			DurationDays:    999,
			Message:         uuid.NewString(),
			Name:            uuid.NewString(),
		},
		{
			ID:              uuid.NewString(),
			AppID:           uuid.NewString(),
			Denomination:    "999999",
			Circulation:     "999999",
			ReleaseByUserID: uuid.NewString(),
			StartAt:         999,
			DurationDays:    999,
			Message:         uuid.NewString(),
			Name:            uuid.NewString(),
		},
	}

	apps := []*npool.FixAmountReq{}
	for key := range appDates {
		apps = append(apps, &npool.FixAmountReq{
			ID:              &appDates[key].ID,
			AppID:           &appDates[key].AppID,
			Denomination:    &appDates[key].Denomination,
			Circulation:     &appDates[key].Circulation,
			ReleaseByUserID: &appDates[key].ReleaseByUserID,
			StartAt:         &appDates[key].StartAt,
			DurationDays:    &appDates[key].DurationDays,
			Message:         &appDates[key].Message,
			Name:            &appDates[key].Name,
		})
	}

	infos, err := CreateFixAmounts(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateFixAmount(t *testing.T) {
	var err error
	info, err = UpdateFixAmount(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func getFixAmount(t *testing.T) {
	var err error
	info, err = GetFixAmount(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getFixAmounts(t *testing.T) {
	infos, total, err := GetFixAmounts(context.Background(),
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

func getFixAmountOnly(t *testing.T) {
	var err error
	info, err = GetFixAmountOnly(context.Background(),
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

func existFixAmount(t *testing.T) {
	exist, err := ExistFixAmount(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existFixAmountConds(t *testing.T) {
	exist, err := ExistFixAmountConds(context.Background(),
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

func deleteFixAmount(t *testing.T) {
	info, err := DeleteFixAmount(context.Background(), info.ID)
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

	t.Run("createFixAmount", createFixAmount)
	t.Run("createFixAmounts", createFixAmounts)
	t.Run("getFixAmount", getFixAmount)
	t.Run("getFixAmounts", getFixAmounts)
	t.Run("getFixAmountOnly", getFixAmountOnly)
	t.Run("updateFixAmount", updateFixAmount)
	t.Run("existFixAmount", existFixAmount)
	t.Run("existFixAmountConds", existFixAmountConds)
	t.Run("delete", deleteFixAmount)
}
