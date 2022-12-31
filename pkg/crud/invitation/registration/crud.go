package registration

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/inspire-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer/invitation/registration"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/registration"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"

	"github.com/google/uuid"
)

func CreateSet(c *ent.RegistrationCreate, in *npool.RegistrationReq) (*ent.RegistrationCreate, error) {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.AppID != nil {
		c.SetAppID(uuid.MustParse(in.GetAppID()))
	}
	if in.InviterID != nil {
		c.SetInviterID(uuid.MustParse(in.GetInviterID()))
	}
	if in.InviteeID != nil {
		c.SetInviteeID(uuid.MustParse(in.GetInviteeID()))
	}

	return c, nil
}

func Create(ctx context.Context, in *npool.RegistrationReq) (*ent.Registration, error) {
	var info *ent.Registration
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
		c, err := CreateSet(cli.Registration.Create(), in)
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

func CreateBulk(ctx context.Context, in []*npool.RegistrationReq) ([]*ent.Registration, error) {
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

	rows := []*ent.Registration{}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.RegistrationCreate, len(in))
		for i, info := range in {
			bulk[i], err = CreateSet(tx.Registration.Create(), info)
			if err != nil {
				return err
			}
		}
		rows, err = tx.Registration.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func UpdateSet(info *ent.Registration, in *npool.RegistrationReq) *ent.RegistrationUpdateOne {
	stm := info.Update()

	if in.InviterID != nil {
		stm = stm.SetInviterID(uuid.MustParse(in.GetInviterID()))
	}

	return stm
}

func Update(ctx context.Context, in *npool.RegistrationReq) (*ent.Registration, error) {
	var info *ent.Registration
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
		info, err = tx.Registration.Query().Where(registration.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return fmt.Errorf("fail query registration: %v", err)
		}

		stm := UpdateSet(info, in)

		info, err = stm.Save(_ctx)
		if err != nil {
			return fmt.Errorf("fail update registration: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update registration: %v", err)
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.Registration, error) {
	var info *ent.Registration
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
		info, err = cli.Registration.Query().Where(registration.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.RegistrationQuery, error) {
	stm := cli.Registration.Query()
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(registration.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid registration field")
		}
	}
	if conds.AppID != nil {
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(registration.AppID(uuid.MustParse(conds.GetAppID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid registration field")
		}
	}
	if conds.InviterID != nil {
		switch conds.GetInviterID().GetOp() {
		case cruder.EQ:
			stm.Where(registration.InviterID(uuid.MustParse(conds.GetInviterID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid registration field")
		}
	}
	if conds.InviteeID != nil {
		switch conds.GetInviteeID().GetOp() {
		case cruder.EQ:
			stm.Where(registration.InviteeID(uuid.MustParse(conds.GetInviteeID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid registration field")
		}
	}
	if len(conds.GetInviterIDs().GetValue()) > 0 {
		ids := []uuid.UUID{}
		for _, id := range conds.GetInviterIDs().GetValue() {
			ids = append(ids, uuid.MustParse(id))
		}
		switch conds.GetInviterIDs().GetOp() {
		case cruder.IN:
			stm.Where(registration.InviterIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid registration field")
		}
	}
	if len(conds.GetInviteeIDs().GetValue()) > 0 {
		ids := []uuid.UUID{}
		for _, id := range conds.GetInviteeIDs().GetValue() {
			ids = append(ids, uuid.MustParse(id))
		}
		switch conds.GetInviteeIDs().GetOp() {
		case cruder.IN:
			stm.Where(registration.InviteeIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid registration field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.Registration, int, error) {
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

	rows := []*ent.Registration{}
	var total int
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(registration.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.Registration, error) {
	var info *ent.Registration
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
		stm, err := SetQueryConds(conds, cli)
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
		stm, err := SetQueryConds(conds, cli)
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
		exist, err = cli.Registration.Query().Where(registration.ID(id)).Exist(_ctx)
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
		stm, err := SetQueryConds(conds, cli)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.Registration, error) {
	var info *ent.Registration
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Delete")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Registration.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
