package event

import (
	"context"
	entevent "github.com/NpoolPlatform/inspire-manager/pkg/db/ent/event"
	constant "github.com/NpoolPlatform/inspire-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/event"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func CreateSet(c *ent.EventCreate, in *npool.EventReq) (*ent.EventCreate, error) {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.AppID != nil {
		c.SetAppID(uuid.MustParse(in.GetAppID()))
	}
	if in.EventType != nil {
		c.SetEventType(in.GetEventType().String())
	}
	if len(in.GetCoupons()) > 0 {
		coupons := []npool.Coupon{}
		for _, coup := range in.GetCoupons() {
			coupons = append(coupons, npool.Coupon{
				ID:         coup.ID,
				CouponType: coup.CouponType,
			})
		}
		c.SetCoupons(coupons)
	}
	if in.Credits != nil {
		c.SetCredits(decimal.RequireFromString(in.GetCredits()))
	}
	if in.CreditsPerUSD != nil {
		c.SetCreditsPerUsd(decimal.RequireFromString(in.GetCreditsPerUSD()))
	}
	if in.MaxConsecutive != nil {
		c.SetMaxConsecutive(in.GetMaxConsecutive())
	}
	if in.GoodID != nil {
		c.SetGoodID(uuid.MustParse(in.GetGoodID()))
	}
	if in.InviterLayers != nil {
		c.SetInviterLayers(in.GetInviterLayers())
	}
	return c, nil
}

func UpdateSet(info *ent.Event, in *npool.EventReq) (*ent.EventUpdateOne, error) {
	u := info.Update()

	if in.Coupons != nil {
		coupons := []npool.Coupon{}
		for _, coup := range in.GetCoupons() {
			coupons = append(coupons, npool.Coupon{
				ID:         coup.ID,
				CouponType: coup.CouponType,
			})
		}
		u.SetCoupons(coupons)
	}
	if in.Credits != nil {
		u.SetCredits(decimal.RequireFromString(in.GetCredits()))
	}
	if in.CreditsPerUSD != nil {
		u.SetCreditsPerUsd(decimal.RequireFromString(in.GetCreditsPerUSD()))
	}
	if in.MaxConsecutive != nil {
		u.SetMaxConsecutive(in.GetMaxConsecutive())
	}
	if in.InviterLayers != nil {
		u.SetInviterLayers(in.GetInviterLayers())
	}

	return u, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.Event, error) {
	var info *ent.Event
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Row")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Event.Query().Where(entevent.ID(id)).Only(_ctx)
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
		exist, err = cli.Event.Query().Where(entevent.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
