package detail

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/inspire-manager/pkg/servicename"
	commontracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer/archivement/detail"

	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/archivementdetail"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/detail"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

func CreateSet(c *ent.ArchivementDetailCreate, in *npool.DetailReq) (*ent.ArchivementDetailCreate, error) { //nolint
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.AppID != nil {
		c.SetAppID(uuid.MustParse(in.GetAppID()))
	}
	if in.UserID != nil {
		c.SetUserID(uuid.MustParse(in.GetUserID()))
	}
	if in.DirectContributorID != nil {
		c.SetDirectContributorID(uuid.MustParse(in.GetDirectContributorID()))
	}
	if in.GoodID != nil {
		c.SetGoodID(uuid.MustParse(in.GetGoodID()))
	}
	if in.OrderID != nil {
		c.SetOrderID(uuid.MustParse(in.GetOrderID()))
	}
	if in.SelfOrder != nil {
		c.SetSelfOrder(in.GetSelfOrder())
	}
	if in.PaymentID != nil {
		c.SetPaymentID(uuid.MustParse(in.GetPaymentID()))
	}
	if in.CoinTypeID != nil {
		c.SetCoinTypeID(uuid.MustParse(in.GetCoinTypeID()))
	}
	if in.PaymentCoinTypeID != nil {
		c.SetPaymentCoinTypeID(uuid.MustParse(in.GetPaymentCoinTypeID()))
	}
	if in.PaymentCoinUSDCurrency != nil {
		currency, err := decimal.NewFromString(in.GetPaymentCoinUSDCurrency())
		if err != nil {
			return nil, err
		}
		c.SetPaymentCoinUsdCurrency(currency)
	}
	if in.Amount != nil {
		amount, err := decimal.NewFromString(in.GetAmount())
		if err != nil {
			return nil, err
		}
		c.SetAmount(amount)
	}
	if in.USDAmount != nil {
		amount, err := decimal.NewFromString(in.GetUSDAmount())
		if err != nil {
			return nil, err
		}
		c.SetUsdAmount(amount)
	}
	if in.Commission != nil {
		amount, err := decimal.NewFromString(in.GetCommission())
		if err != nil {
			return nil, err
		}
		c.SetCommission(amount)
	}
	if in.Units != nil {
		units, err := decimal.NewFromString(in.GetUnits())
		if err != nil {
			return nil, err
		}
		c.SetUnitsV1(units)
	}
	if in.CreatedAt != nil {
		c.SetCreatedAt(in.GetCreatedAt())
	}

	return c, nil
}

func Create(ctx context.Context, in *npool.DetailReq) (*ent.ArchivementDetail, error) {
	var info *ent.ArchivementDetail
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c, err := CreateSet(cli.ArchivementDetail.Create(), in)
		if err != nil {
			return err
		}

		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.DetailReq) ([]*ent.ArchivementDetail, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CreateBulk")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceMany(span, in)

	rows := []*ent.ArchivementDetail{}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.ArchivementDetailCreate, len(in))
		for i, info := range in {
			bulk[i], err = CreateSet(tx.ArchivementDetail.Create(), info)
			if err != nil {
				return err
			}
		}
		rows, err = tx.ArchivementDetail.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.ArchivementDetail, error) {
	var info *ent.ArchivementDetail
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "Row")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.ArchivementDetail.Query().Where(archivementdetail.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.ArchivementDetailQuery, error) { //nolint
	stm := cli.ArchivementDetail.Query()
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(archivementdetail.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.AppID != nil {
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(archivementdetail.AppID(uuid.MustParse(conds.GetAppID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.UserID != nil {
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(archivementdetail.UserID(uuid.MustParse(conds.GetUserID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if len(conds.GetUserIDs().GetValue()) > 0 {
		ids := []uuid.UUID{}
		for _, id := range conds.GetUserIDs().GetValue() {
			ids = append(ids, uuid.MustParse(id))
		}
		switch conds.GetUserIDs().GetOp() {
		case cruder.IN:
			stm.Where(archivementdetail.UserIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.DirectContributorID != nil {
		switch conds.GetDirectContributorID().GetOp() {
		case cruder.EQ:
			stm.Where(archivementdetail.DirectContributorID(uuid.MustParse(conds.GetDirectContributorID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.GoodID != nil {
		switch conds.GetGoodID().GetOp() {
		case cruder.EQ:
			stm.Where(archivementdetail.GoodID(uuid.MustParse(conds.GetGoodID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.OrderID != nil {
		switch conds.GetOrderID().GetOp() {
		case cruder.EQ:
			stm.Where(archivementdetail.OrderID(uuid.MustParse(conds.GetOrderID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.SelfOrder != nil {
		switch conds.GetSelfOrder().GetOp() {
		case cruder.EQ:
			stm.Where(archivementdetail.SelfOrder(conds.GetSelfOrder().GetValue()))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.PaymentID != nil {
		switch conds.GetPaymentID().GetOp() {
		case cruder.EQ:
			stm.Where(archivementdetail.PaymentID(uuid.MustParse(conds.GetPaymentID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.CoinTypeID != nil {
		switch conds.GetCoinTypeID().GetOp() {
		case cruder.EQ:
			stm.Where(archivementdetail.CoinTypeID(uuid.MustParse(conds.GetCoinTypeID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.PaymentCoinTypeID != nil {
		switch conds.GetPaymentCoinTypeID().GetOp() {
		case cruder.EQ:
			stm.Where(archivementdetail.PaymentCoinTypeID(uuid.MustParse(conds.GetPaymentCoinTypeID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.PaymentCoinUSDCurrency != nil {
		currency, err := decimal.NewFromString(conds.GetPaymentCoinUSDCurrency().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetPaymentCoinUSDCurrency().GetOp() {
		case cruder.LT:
			stm.Where(archivementdetail.PaymentCoinUsdCurrencyLT(currency))
		case cruder.GT:
			stm.Where(archivementdetail.PaymentCoinUsdCurrencyGT(currency))
		case cruder.EQ:
			stm.Where(archivementdetail.PaymentCoinUsdCurrencyEQ(currency))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.Amount != nil {
		amount, err := decimal.NewFromString(conds.GetAmount().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetAmount().GetOp() {
		case cruder.LT:
			stm.Where(archivementdetail.AmountLT(amount))
		case cruder.GT:
			stm.Where(archivementdetail.AmountGT(amount))
		case cruder.EQ:
			stm.Where(archivementdetail.AmountEQ(amount))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.USDAmount != nil {
		amount, err := decimal.NewFromString(conds.GetUSDAmount().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetUSDAmount().GetOp() {
		case cruder.LT:
			stm.Where(archivementdetail.UsdAmountLT(amount))
		case cruder.GT:
			stm.Where(archivementdetail.UsdAmountGT(amount))
		case cruder.EQ:
			stm.Where(archivementdetail.UsdAmountEQ(amount))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.Commission != nil {
		amount, err := decimal.NewFromString(conds.GetCommission().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetCommission().GetOp() {
		case cruder.LT:
			stm.Where(archivementdetail.CommissionLT(amount))
		case cruder.GT:
			stm.Where(archivementdetail.CommissionGT(amount))
		case cruder.EQ:
			stm.Where(archivementdetail.CommissionEQ(amount))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.Units != nil {
		units, err := decimal.NewFromString(conds.GetUnits().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetUnits().GetOp() {
		case cruder.LT:
			stm.Where(archivementdetail.UnitsV1LT(units))
		case cruder.GT:
			stm.Where(archivementdetail.UnitsV1GT(units))
		case cruder.EQ:
			stm.Where(archivementdetail.UnitsV1EQ(units))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.ArchivementDetail, int, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "Rows")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)
	span = commontracer.TraceOffsetLimit(span, offset, limit)

	rows := []*ent.ArchivementDetail{}
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
			Order(ent.Desc(archivementdetail.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.ArchivementDetail, error) {
	var info *ent.ArchivementDetail
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "RowOnly")
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
			if ent.IsNotFound(err) {
				return nil
			}
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

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "Count")
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

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "Exist")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.ArchivementDetail.Query().Where(archivementdetail.ID(id)).Exist(_ctx)
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

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "ExistConds")
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.ArchivementDetail, error) {
	var info *ent.ArchivementDetail
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "Delete")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.ArchivementDetail.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
