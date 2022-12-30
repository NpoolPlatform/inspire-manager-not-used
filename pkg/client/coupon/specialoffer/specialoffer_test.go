package specialoffer

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

	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/specialoffer"

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

var appDate = npool.SpecialOffer{
	ID:               uuid.NewString(),
	AppID:            uuid.NewString(),
	UserID:           uuid.NewString(),
	Amount:           "99999",
	ReleasedByUserID: uuid.NewString(),
	StartAt:          999,
	DurationDays:     999,
	Message:          uuid.NewString(),
}

var (
	appInfo = npool.SpecialOfferReq{
		ID:               &appDate.ID,
		AppID:            &appDate.AppID,
		UserID:           &appDate.UserID,
		Amount:           &appDate.Amount,
		ReleasedByUserID: &appDate.ReleasedByUserID,
		StartAt:          &appDate.StartAt,
		DurationDays:     &appDate.DurationDays,
		Message:          &appDate.Message,
	}
)

var info *npool.SpecialOffer

func createSpecialOffer(t *testing.T) {
	var err error
	info, err = CreateSpecialOffer(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.CreatedAt = info.CreatedAt
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func createSpecialOffers(t *testing.T) {
	appDates := []npool.SpecialOffer{
		{
			ID:               uuid.NewString(),
			AppID:            uuid.NewString(),
			UserID:           uuid.NewString(),
			Amount:           "99999",
			ReleasedByUserID: uuid.NewString(),
			StartAt:          999,
			DurationDays:     999,
			Message:          uuid.NewString(),
		},
		{
			ID:               uuid.NewString(),
			AppID:            uuid.NewString(),
			UserID:           uuid.NewString(),
			Amount:           "99999",
			ReleasedByUserID: uuid.NewString(),
			StartAt:          999,
			DurationDays:     999,
			Message:          uuid.NewString(),
		},
	}

	apps := []*npool.SpecialOfferReq{}
	for key := range appDates {
		apps = append(apps, &npool.SpecialOfferReq{
			ID:               &appDates[key].ID,
			AppID:            &appDates[key].AppID,
			UserID:           &appDates[key].UserID,
			Amount:           &appDates[key].Amount,
			ReleasedByUserID: &appDates[key].ReleasedByUserID,
			StartAt:          &appDates[key].StartAt,
			DurationDays:     &appDates[key].DurationDays,
			Message:          &appDates[key].Message,
		})
	}

	infos, err := CreateSpecialOffers(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateSpecialOffer(t *testing.T) {
	var err error
	info, err = UpdateSpecialOffer(context.Background(), &appInfo)
	fmt.Println(err)
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func getSpecialOffer(t *testing.T) {
	var err error
	info, err = GetSpecialOffer(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getSpecialOffers(t *testing.T) {
	infos, total, err := GetSpecialOffers(context.Background(),
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

func getSpecialOfferOnly(t *testing.T) {
	var err error
	info, err = GetSpecialOfferOnly(context.Background(),
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

func existSpecialOffer(t *testing.T) {
	exist, err := ExistSpecialOffer(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existSpecialOfferConds(t *testing.T) {
	exist, err := ExistSpecialOfferConds(context.Background(),
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

func deleteSpecialOffer(t *testing.T) {
	info, err := DeleteSpecialOffer(context.Background(), info.ID)
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

	t.Run("createSpecialOffer", createSpecialOffer)
	t.Run("createSpecialOffers", createSpecialOffers)
	t.Run("getSpecialOffer", getSpecialOffer)
	t.Run("getSpecialOffers", getSpecialOffers)
	t.Run("getSpecialOfferOnly", getSpecialOfferOnly)
	t.Run("updateSpecialOffer", updateSpecialOffer)
	t.Run("existSpecialOffer", existSpecialOffer)
	t.Run("existSpecialOfferConds", existSpecialOfferConds)
	t.Run("delete", deleteSpecialOffer)
}
