package allocated

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

var appDate = npool.Allocated{
	ID:       uuid.NewString(),
	AppID:    uuid.NewString(),
	UserID:   uuid.NewString(),
	Type:     npool.CouponType_Discount,
	CouponID: uuid.NewString(),
}

var (
	appInfo = npool.AllocatedReq{
		ID:       &appDate.ID,
		AppID:    &appDate.AppID,
		UserID:   &appDate.UserID,
		Type:     &appDate.Type,
		CouponID: &appDate.CouponID,
	}
)

var info *npool.Allocated

func createAllocated(t *testing.T) {
	var err error
	info, err = CreateAllocated(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.CreatedAt = info.CreatedAt
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func createAllocateds(t *testing.T) {
	appDates := []npool.Allocated{
		{
			ID:       uuid.NewString(),
			AppID:    uuid.NewString(),
			UserID:   uuid.NewString(),
			Type:     npool.CouponType_Discount,
			CouponID: uuid.NewString(),
		},
		{
			ID:       uuid.NewString(),
			AppID:    uuid.NewString(),
			UserID:   uuid.NewString(),
			Type:     npool.CouponType_Discount,
			CouponID: uuid.NewString(),
		},
	}

	apps := []*npool.AllocatedReq{}
	for key := range appDates {
		apps = append(apps, &npool.AllocatedReq{
			ID:       &appDates[key].ID,
			AppID:    &appDates[key].AppID,
			UserID:   &appDates[key].UserID,
			Type:     &appDates[key].Type,
			CouponID: &appDates[key].CouponID,
		})
	}

	infos, err := CreateAllocateds(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func getAllocated(t *testing.T) {
	var err error
	info, err = GetAllocated(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getAllocateds(t *testing.T) {
	infos, total, err := GetAllocateds(context.Background(),
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

func getAllocatedOnly(t *testing.T) {
	var err error
	info, err = GetAllocatedOnly(context.Background(),
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

func existAllocated(t *testing.T) {
	exist, err := ExistAllocated(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existAllocatedConds(t *testing.T) {
	exist, err := ExistAllocatedConds(context.Background(),
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

func deleteAllocated(t *testing.T) {
	info, err := DeleteAllocated(context.Background(), info.ID)
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

	t.Run("createAllocated", createAllocated)
	t.Run("createAllocateds", createAllocateds)
	t.Run("getAllocated", getAllocated)
	t.Run("getAllocateds", getAllocateds)
	t.Run("getAllocatedOnly", getAllocatedOnly)
	t.Run("existAllocated", existAllocated)
	t.Run("existAllocatedConds", existAllocatedConds)
	t.Run("delete", deleteAllocated)
}
