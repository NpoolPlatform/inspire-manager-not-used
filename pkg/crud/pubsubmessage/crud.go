package pubsubmessage

import (
	"context"

	entpubsubmessgae "github.com/NpoolPlatform/inspire-manager/pkg/db/ent/pubsubmessage"
	constant "github.com/NpoolPlatform/inspire-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
	"github.com/google/uuid"
)

func CreateSet(
	c *ent.PubsubMessageCreate,
	uniqueID uuid.UUID,
	messageID string,
	state string,
	respondToID *uuid.UUID,
) (*ent.PubsubMessageCreate, error) {
	c.SetID(uniqueID)
	c.SetMessageID(messageID)
	c.SetState(state)
	respondToID1 := uuid.UUID{}
	if respondToID != nil {
		respondToID1 = *respondToID
	}
	c.SetResponseToID(respondToID1)
	return c, nil
}

func UpdateSet(
	info *ent.PubsubMessage,
	state string,
) (*ent.PubsubMessageUpdateOne, error) {
	u := info.Update()
	u.SetState(state)
	return u, nil
}

func Row(ctx context.Context, uniqueID uuid.UUID) (*ent.PubsubMessage, error) {
	var info *ent.PubsubMessage
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Row")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, uniqueID.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.PubsubMessage.Query().Where(entpubsubmessgae.ID(uniqueID)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Exist(ctx context.Context, id uuid.UUID) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Exist")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.PubsubMessage.Query().Where(entpubsubmessgae.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}