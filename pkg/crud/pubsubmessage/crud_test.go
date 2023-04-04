package pubsubmessage

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/pubsubmessage"

	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
	testinit "github.com/NpoolPlatform/inspire-manager/pkg/testinit"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

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

var data = ent.PubsubMessage{
	ID:           uuid.New(),
	MessageID:    uuid.NewString(),
	State:        basetypes.MessageState_Processing.String(),
	ResponseToID: uuid.New(),
}

var info *ent.PubsubMessage

func create(t *testing.T) {
	err := db.WithClient(context.Background(), func(ctx context.Context, cli *ent.Client) error {
		c, err := CreateSet(
			cli.PubsubMessage.Create(),
			data.ID,
			data.MessageID,
			data.State,
			&data.ResponseToID,
		)
		if err != nil {
			return err
		}
		info, err = c.Save(ctx)
		return err
	})
	if assert.Nil(t, err) {
		data.UpdatedAt = info.UpdatedAt
		data.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), data.String())
	}
}

func update(t *testing.T) {
	var err error
	err = db.WithClient(context.Background(), func(ctx context.Context, cli *ent.Client) error {
		info, err = cli.
			PubsubMessage.
			Query().
			Where(
				pubsubmessage.ID(data.ID),
			).
			ForUpdate().
			Only(ctx)
		u, err := UpdateSet(info, basetypes.MessageState_Success.String())
		if err != nil {
			return err
		}
		info, err = u.Save(ctx)
		return err
	})
	if assert.Nil(t, err) {
		data.UpdatedAt = info.UpdatedAt
		data.CreatedAt = info.CreatedAt
		data.State = basetypes.MessageState_Success.String()
		assert.Equal(t, info.String(), data.String())
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), data.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), data.String())
	}
}

func exist(t *testing.T) {
	exist, err := Exist(context.Background(), data.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func TestDetail(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("update", update)
	t.Run("row", row)
	t.Run("exist", exist)
}
