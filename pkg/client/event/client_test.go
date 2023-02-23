package event

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

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/event"

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

var ret = npool.Event{
	ID:             uuid.NewString(),
	AppID:          uuid.NewString(),
	EventType:      basetypes.UsedFor_Signup,
	CouponIDs:      []string{uuid.New().String(), uuid.New().String()},
	Credits:        "1.798",
	CreditsPerUSD:  "1.789",
	MaxConsecutive: 1,
	GoodID:         uuid.UUID{}.String(),
}

var (
	req = npool.EventReq{
		ID:            &ret.ID,
		AppID:         &ret.AppID,
		EventType:     &ret.EventType,
		CouponIDs:     ret.CouponIDs,
		Credits:       &ret.Credits,
		CreditsPerUSD: &ret.CreditsPerUSD,
	}
)

var info *npool.Event

func createEvent(t *testing.T) {
	var err error
	info, err = CreateEvent(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func createEvents(t *testing.T) {
	rets := []npool.Event{
		{
			ID:            uuid.NewString(),
			AppID:         uuid.NewString(),
			EventType:     basetypes.UsedFor_Signin,
			CouponIDs:     []string{uuid.New().String(), uuid.New().String()},
			Credits:       "1.7981",
			CreditsPerUSD: "1.7892",
			GoodID:        uuid.UUID{}.String(),
		},
		{
			ID:            uuid.NewString(),
			AppID:         uuid.NewString(),
			EventType:     basetypes.UsedFor_Signin,
			CouponIDs:     []string{uuid.New().String(), uuid.New().String()},
			Credits:       "1.7983",
			CreditsPerUSD: "1.7894",
			GoodID:        uuid.UUID{}.String(),
		},
	}

	apps := []*npool.EventReq{}
	for key := range rets {
		apps = append(apps, &npool.EventReq{
			ID:            &rets[key].ID,
			AppID:         &rets[key].AppID,
			EventType:     &rets[key].EventType,
			CouponIDs:     rets[key].CouponIDs,
			Credits:       &rets[key].Credits,
			CreditsPerUSD: &rets[key].CreditsPerUSD,
		})
	}

	infos, err := CreateEvents(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func getEvent(t *testing.T) {
	var err error
	info, err = GetEvent(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getEvents(t *testing.T) {
	infos, total, err := GetEvents(context.Background(),
		&npool.Conds{
			ID: &basetypes.StringVal{Value: info.ID, Op: cruder.EQ},
		}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &ret)
	}
}

func getEventOnly(t *testing.T) {
	var err error
	info, err = GetEventOnly(context.Background(),
		&npool.Conds{
			ID: &basetypes.StringVal{Value: info.ID, Op: cruder.EQ},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func existEvent(t *testing.T) {
	exist, err := ExistEvent(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existEventConds(t *testing.T) {
	exist, err := ExistEventConds(context.Background(),
		&npool.Conds{
			ID: &basetypes.StringVal{Value: info.ID, Op: cruder.EQ},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteEvent(t *testing.T) {
	info, err := DeleteEvent(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
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

	t.Run("createEvent", createEvent)
	t.Run("createEvents", createEvents)
	t.Run("getEvent", getEvent)
	t.Run("getEvents", getEvents)
	t.Run("getEventOnly", getEventOnly)
	t.Run("existEvent", existEvent)
	t.Run("existEventConds", existEventConds)
	t.Run("delete", deleteEvent)
}
