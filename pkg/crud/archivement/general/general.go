package general

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/inspire-manager/pkg/servicename"
	commontracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer/archivement/general"

	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/archivementgeneral"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/general"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

func CreateSet(c *ent.ArchivementGeneralCreate, in *npool.GeneralReq) *ent.ArchivementGeneralCreate {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.AppID != nil {
		c.SetAppID(uuid.MustParse(in.GetAppID()))
	}
	if in.UserID != nil {
		c.SetUserID(uuid.MustParse(in.GetUserID()))
	}
	if in.GoodID != nil {
		c.SetGoodID(uuid.MustParse(in.GetGoodID()))
	}
	if in.CoinTypeID != nil {
		c.SetCoinTypeID(uuid.MustParse(in.GetCoinTypeID()))
	}

	c.SetTotalAmount(decimal.NewFromInt(0))
	c.SetSelfAmount(decimal.NewFromInt(0))
	c.SetTotalUnitsV1(decimal.NewFromInt(0))
	c.SetSelfUnitsV1(decimal.NewFromInt(0))
	c.SetTotalCommission(decimal.NewFromInt(0))
	c.SetSelfCommission(decimal.NewFromInt(0))

	return c
}

func Create(ctx context.Context, in *npool.GeneralReq) (*ent.ArchivementGeneral, error) {
	var info *ent.ArchivementGeneral
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
		c := CreateSet(cli.ArchivementGeneral.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.GeneralReq) ([]*ent.ArchivementGeneral, error) {
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

	rows := []*ent.ArchivementGeneral{}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.ArchivementGeneralCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.ArchivementGeneral.Create(), info)
		}
		rows, err = tx.ArchivementGeneral.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

//nolint:gocyclo
func UpdateSet(info *ent.ArchivementGeneral, in *npool.GeneralReq) (u *ent.ArchivementGeneralUpdateOne, err error) {
	totalAmount := decimal.NewFromInt(0)
	if in.TotalAmount != nil {
		totalAmount, err = decimal.NewFromString(in.GetTotalAmount())
		if err != nil {
			return nil, err
		}
	}

	selfAmount := decimal.NewFromInt(0)
	if in.SelfAmount != nil {
		selfAmount, err = decimal.NewFromString(in.GetSelfAmount())
		if err != nil {
			return nil, err
		}
	}

	totalCommission := decimal.NewFromInt(0)
	if in.TotalCommission != nil {
		totalCommission, err = decimal.NewFromString(in.GetTotalCommission())
		if err != nil {
			return nil, err
		}
	}

	selfCommission := decimal.NewFromInt(0)
	if in.SelfCommission != nil {
		selfCommission, err = decimal.NewFromString(in.GetSelfCommission())
		if err != nil {
			return nil, err
		}
	}

	stm := info.Update()

	if in.TotalAmount != nil {
		totalAmount = totalAmount.Add(info.TotalAmount)
		stm = stm.SetTotalAmount(totalAmount)
	}
	if in.SelfAmount != nil {
		selfAmount = selfAmount.Add(info.SelfAmount)
		stm = stm.SetSelfAmount(selfAmount)
	}
	if in.TotalUnits != nil {
		totalUnits, err := decimal.NewFromString(in.GetTotalUnits())
		if err != nil {
			return nil, err
		}
		totalUnits = totalUnits.Add(info.TotalUnitsV1)
		stm = stm.SetTotalUnitsV1(totalUnits)
	}
	if in.SelfUnits != nil {
		selfUnits, err := decimal.NewFromString(in.GetSelfUnits())
		if err != nil {
			return nil, err
		}
		selfUnits = selfUnits.Add(info.SelfUnitsV1)
		stm = stm.SetSelfUnitsV1(selfUnits)
	}
	if in.TotalCommission != nil {
		totalCommission = totalCommission.Add(info.TotalCommission)
		stm = stm.SetTotalCommission(totalCommission)
	}
	if in.SelfCommission != nil {
		selfCommission = selfCommission.Add(info.SelfCommission)
		stm = stm.SetSelfCommission(selfCommission)
	}

	return stm, nil
}

func AddFields(ctx context.Context, in *npool.GeneralReq) (*ent.ArchivementGeneral, error) {
	var info *ent.ArchivementGeneral
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

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err = tx.ArchivementGeneral.Query().Where(archivementgeneral.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return fmt.Errorf("fail query general: %v", err)
		}

		stm, err := UpdateSet(info, in)
		if err != nil {
			return err
		}

		info, err = stm.Save(_ctx)
		if err != nil {
			return fmt.Errorf("fail update general: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update general: %v", err)
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.ArchivementGeneral, error) {
	var info *ent.ArchivementGeneral
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
		info, err = cli.ArchivementGeneral.Query().Where(archivementgeneral.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.ArchivementGeneralQuery, error) { //nolint
	stm := cli.ArchivementGeneral.Query()
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(archivementgeneral.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.AppID != nil {
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(archivementgeneral.AppID(uuid.MustParse(conds.GetAppID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.UserID != nil {
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(archivementgeneral.UserID(uuid.MustParse(conds.GetUserID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.UserIDs != nil {
		switch conds.GetUserIDs().GetOp() {
		case cruder.IN:
			users := []uuid.UUID{}
			for _, user := range conds.GetUserIDs().GetValue() {
				users = append(users, uuid.MustParse(user))
			}
			stm.Where(archivementgeneral.UserIDIn(users...))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.GoodID != nil {
		switch conds.GetGoodID().GetOp() {
		case cruder.EQ:
			stm.Where(archivementgeneral.GoodID(uuid.MustParse(conds.GetGoodID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.CoinTypeID != nil {
		switch conds.GetCoinTypeID().GetOp() {
		case cruder.EQ:
			stm.Where(archivementgeneral.CoinTypeID(uuid.MustParse(conds.GetCoinTypeID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.TotalAmount != nil {
		amount, err := decimal.NewFromString(conds.GetTotalAmount().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetTotalAmount().GetOp() {
		case cruder.LT:
			stm.Where(archivementgeneral.TotalAmountLT(amount))
		case cruder.GT:
			stm.Where(archivementgeneral.TotalAmountGT(amount))
		case cruder.EQ:
			stm.Where(archivementgeneral.TotalAmountEQ(amount))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.SelfAmount != nil {
		amount, err := decimal.NewFromString(conds.GetSelfAmount().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetSelfAmount().GetOp() {
		case cruder.LT:
			stm.Where(archivementgeneral.SelfAmountLT(amount))
		case cruder.GT:
			stm.Where(archivementgeneral.SelfAmountGT(amount))
		case cruder.EQ:
			stm.Where(archivementgeneral.SelfAmountEQ(amount))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.TotalUnits != nil {
		totalUnits, err := decimal.NewFromString(conds.GetTotalUnits().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetTotalUnits().GetOp() {
		case cruder.LT:
			stm.Where(archivementgeneral.TotalUnitsV1LT(totalUnits))
		case cruder.GT:
			stm.Where(archivementgeneral.TotalUnitsV1GT(totalUnits))
		case cruder.EQ:
			stm.Where(archivementgeneral.TotalUnitsV1EQ(totalUnits))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.SelfUnits != nil {
		selfUnits, err := decimal.NewFromString(conds.GetSelfUnits().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetSelfUnits().GetOp() {
		case cruder.LT:
			stm.Where(archivementgeneral.SelfUnitsV1LT(selfUnits))
		case cruder.GT:
			stm.Where(archivementgeneral.SelfUnitsV1GT(selfUnits))
		case cruder.EQ:
			stm.Where(archivementgeneral.SelfUnitsV1EQ(selfUnits))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.TotalCommission != nil {
		amount, err := decimal.NewFromString(conds.GetTotalCommission().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetTotalCommission().GetOp() {
		case cruder.LT:
			stm.Where(archivementgeneral.TotalCommissionLT(amount))
		case cruder.GT:
			stm.Where(archivementgeneral.TotalCommissionGT(amount))
		case cruder.EQ:
			stm.Where(archivementgeneral.TotalCommissionEQ(amount))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.SelfCommission != nil {
		amount, err := decimal.NewFromString(conds.GetSelfCommission().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetSelfCommission().GetOp() {
		case cruder.LT:
			stm.Where(archivementgeneral.SelfCommissionLT(amount))
		case cruder.GT:
			stm.Where(archivementgeneral.SelfCommissionGT(amount))
		case cruder.EQ:
			stm.Where(archivementgeneral.SelfCommissionEQ(amount))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.ArchivementGeneral, int, error) {
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

	rows := []*ent.ArchivementGeneral{}
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
			Order(ent.Desc(archivementgeneral.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.ArchivementGeneral, error) {
	var info *ent.ArchivementGeneral
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
		exist, err = cli.ArchivementGeneral.Query().Where(archivementgeneral.ID(id)).Exist(_ctx)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.ArchivementGeneral, error) {
	var info *ent.ArchivementGeneral
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
		info, err = cli.ArchivementGeneral.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
