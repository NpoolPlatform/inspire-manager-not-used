package invitationcode

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
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/invitationcode"

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

var ret = ent.InvitationCode{
	ID:             uuid.New(),
	AppID:          uuid.New(),
	UserID:         uuid.New(),
	InvitationCode: uuid.NewString(),
}

var (
	id     = ret.ID.String()
	appID  = ret.AppID.String()
	userID = ret.UserID.String()

	req = npool.InvitationCodeReq{
		ID:             &id,
		AppID:          &appID,
		UserID:         &userID,
		InvitationCode: &ret.InvitationCode,
	}
)

var info *ent.InvitationCode

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), ret.String())
	}
}

func createBulk(t *testing.T) {
	entities := []*ent.InvitationCode{
		{
			ID:             uuid.New(),
			AppID:          uuid.New(),
			UserID:         uuid.New(),
			InvitationCode: uuid.NewString(),
		},
		{
			ID:             uuid.New(),
			AppID:          uuid.New(),
			UserID:         uuid.New(),
			InvitationCode: uuid.NewString(),
		},
	}

	reqs := []*npool.InvitationCodeReq{}
	for _, _ret := range entities {
		_id := _ret.ID.String()
		_appID := _ret.AppID.String()
		_userID := _ret.UserID.String()

		reqs = append(reqs, &npool.InvitationCodeReq{
			ID:             &_id,
			AppID:          &_appID,
			UserID:         &_userID,
			InvitationCode: &_ret.InvitationCode,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func update(t *testing.T) {
	confirmed := true

	ret.Confirmed = confirmed
	req.Confirmed = &confirmed

	info, err := Update(context.Background(), &req)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), ret.String())
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), ret.ID)
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
	var err error
	info, err = RowOnly(context.Background(),
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
		assert.Equal(t, info.String(), ret.String())
	}
}

func TestInvitationCode(t *testing.T) {
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
