package specialoffer

import (
	"context"
	"fmt"

	"time"

	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/couponspecialoffer"
	tracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer/coupon/specialoffer"
	"github.com/shopspring/decimal"

	constant "github.com/NpoolPlatform/inspire-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/specialoffer"

	"github.com/google/uuid"
)

func CreateSet(c *ent.CouponSpecialOfferCreate, in *npool.SpecialOfferReq) (*ent.CouponSpecialOfferCreate, error) {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.AppID != nil {
		c.SetAppID(uuid.MustParse(in.GetAppID()))
	}
	if in.UserID != nil {
		c.SetUserID(uuid.MustParse(in.GetUserID()))
	}
	if in.Amount != nil {
		val, err := decimal.NewFromString(in.GetAmount())
		if err != nil {
			return nil, err
		}
		c.SetAmount(val)
	}
	if in.ReleaseByUserID != nil {
		c.SetReleaseByUserID(uuid.MustParse(in.GetReleaseByUserID()))
	}
	if in.StartAt != nil {
		c.SetStartAt(in.GetStartAt())
	}
	if in.DurationDays != nil {
		c.SetDurationDays(in.GetDurationDays())
	}
	if in.Message != nil {
		c.SetMessage(in.GetMessage())
	}
	return c, nil
}

func Create(ctx context.Context, in *npool.SpecialOfferReq) (*ent.CouponSpecialOffer, error) {
	var info *ent.CouponSpecialOffer
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := cli.CouponSpecialOffer.Create()
		stm, err := CreateSet(c, in)
		if err != nil {
			return err
		}
		info, err = stm.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.SpecialOfferReq) ([]*ent.CouponSpecialOffer, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBulk")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceMany(span, in)

	rows := []*ent.CouponSpecialOffer{}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.CouponSpecialOfferCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.CouponSpecialOffer.Create()
			bulk[i], err = CreateSet(bulk[i], info)
			if err != nil {
				return err
			}
		}
		rows, err = tx.CouponSpecialOffer.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func UpdateSet(u *ent.CouponSpecialOfferUpdateOne, in *npool.SpecialOfferReq) (*ent.CouponSpecialOfferUpdateOne, error) {
	if in.Amount != nil {
		val, err := decimal.NewFromString(in.GetAmount())
		if err != nil {
			return nil, err
		}
		u.SetAmount(val)
	}
	if in.StartAt != nil {
		u.SetStartAt(in.GetStartAt())
	}
	if in.DurationDays != nil {
		u.SetDurationDays(in.GetDurationDays())
	}
	if in.Message != nil {
		u.SetMessage(in.GetMessage())
	}
	return u, nil
}

func Update(ctx context.Context, in *npool.SpecialOfferReq) (*ent.CouponSpecialOffer, error) {
	var info *ent.CouponSpecialOffer
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Update")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err = tx.CouponSpecialOffer.Query().Where(couponspecialoffer.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return err
		}

		stm, err := UpdateSet(info.Update(), in)
		if err != nil {
			return err
		}

		info, err = stm.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.CouponSpecialOffer, error) {
	var info *ent.CouponSpecialOffer
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
		info, err = cli.CouponSpecialOffer.Query().Where(couponspecialoffer.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.CouponSpecialOfferQuery, error) {
	stm := cli.CouponSpecialOffer.Query()
	if conds == nil {
		return stm, nil
	}
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(couponspecialoffer.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid couponspecialoffer field")
		}
	}
	if conds.AppID != nil {
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(couponspecialoffer.AppID(uuid.MustParse(conds.GetAppID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid couponspecialoffer field")
		}
	}
	if conds.UserID != nil {
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(couponspecialoffer.UserID(uuid.MustParse(conds.GetUserID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid couponspecialoffer field")
		}
	}
	if conds.StartAt != nil {
		switch conds.GetStartAt().GetOp() {
		case cruder.EQ:
			stm.Where(couponspecialoffer.StartAt(conds.GetStartAt().GetValue()))
		default:
			return nil, fmt.Errorf("invalid couponspecialoffer field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.CouponSpecialOffer, int, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Rows")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)
	span = commontracer.TraceOffsetLimit(span, offset, limit)

	rows := []*ent.CouponSpecialOffer{}
	var total int
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(couponspecialoffer.FieldUpdatedAt)).
			Limit(limit).
			All(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.CouponSpecialOffer, error) {
	var info *ent.CouponSpecialOffer
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "RowOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Count(ctx context.Context, conds *npool.Conds) (uint32, error) {
	var err error
	var total int

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Count")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return uint32(total), nil
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
		exist, err = cli.CouponSpecialOffer.Query().Where(couponspecialoffer.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func Delete(ctx context.Context, id string) (*ent.CouponSpecialOffer, error) {
	var info *ent.CouponSpecialOffer
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Delete")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.CouponSpecialOffer.UpdateOneID(uuid.MustParse(id)).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
