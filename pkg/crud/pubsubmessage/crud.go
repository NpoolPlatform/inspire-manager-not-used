package pubsubmessage

import (
	"context"

	entpubsubmessgae "github.com/NpoolPlatform/inspire-manager/pkg/db/ent/pubsubmessgae"
	constant "github.com/NpoolPlatform/inspire-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/go-service-framework/pkg/pubsub"
	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
	"github.com/google/uuid"
)

func CreateSet(
	c *ent.PubsubMessgaeCreate,
	uniqueID uuid.UUID,
	messageID, sender string,
	body []byte,
) (*ent.PubsubMessgaeCreate, error) {
	c.SetUniqueID(uniqueID)
	c.SetMessageID(messageID)
	c.SetSender(sender)
	c.SetBody(body)
	return c, nil
}

func UpdateSet(
	info *ent.PubsubMessgae,
	in *pubsub.Message,
	uniqueID uuid.UUID,
	messageID, sender string,
	body []byte,
) (*ent.PubsubMessgaeUpdateOne, error) {
	u := info.Update()
	u.SetUniqueID(uniqueID)
	u.SetMessageID(messageID)
	u.SetSender(sender)
	u.SetBody(body)
	return u, nil
}

func Row(ctx context.Context, uniqueID uuid.UUID) (*ent.PubsubMessgae, error) {
	var info *ent.PubsubMessgae
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
		info, err = cli.PubsubMessgae.Query().Where(entpubsubmessgae.UniqueID(uniqueID)).Only(_ctx)
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
		exist, err = cli.PubsubMessgae.Query().Where(entpubsubmessgae.UniqueID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
