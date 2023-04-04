//nolint:dupl
package goodorderpercent

import (
	"context"
	"fmt"

	converter "github.com/NpoolPlatform/inspire-manager/pkg/converter/commission/goodorderpercent"
	crud "github.com/NpoolPlatform/inspire-manager/pkg/crud/commission/goodorderpercent"
	commontracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer/commission/goodorderpercent"

	"github.com/NpoolPlatform/inspire-manager/pkg/servicename"

	"github.com/shopspring/decimal"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission/goodorderpercent"

	"github.com/google/uuid"
)

func ValidateCreate(info *npool.OrderPercentReq) error {
	if info.ID != nil {
		if _, err := uuid.Parse(info.GetID()); err != nil {
			logger.Sugar().Errorw("validate", "ID", info.ID, "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("ID is invalid: %v", err))
		}
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", info.AppID, "error", err)
		return err
	}

	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		logger.Sugar().Errorw("validate", "UserID", info.UserID, "error", err)
		return err
	}
	if _, err := uuid.Parse(info.GetGoodID()); err != nil {
		logger.Sugar().Errorw("validate", "GoodID", info.GoodID, "error", err)
		return err
	}

	if _, err := decimal.NewFromString(info.GetPercent()); err != nil {
		logger.Sugar().Errorw("validate", "Percent", info.GetPercent())
		return err
	}
	if info.GetEndAt() <= info.GetStartAt() {
		logger.Sugar().Errorw("validate", "StartAt", info.GetStartAt(), "EndAt", info.GetEndAt())
		return fmt.Errorf("endat <= startat")
	}

	return nil
}

func (s *Server) CreateOrderPercent(
	ctx context.Context,
	in *npool.CreateOrderPercentRequest,
) (
	*npool.CreateOrderPercentResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CreateOrderPercent")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	err = ValidateCreate(in.GetInfo())
	if err != nil {
		return &npool.CreateOrderPercentResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "goodorderpercent", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create goodorderpercent: %v", err.Error())
		return &npool.CreateOrderPercentResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateOrderPercentResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateOrderPercents(
	ctx context.Context,
	in *npool.CreateOrderPercentsRequest,
) (
	*npool.CreateOrderPercentsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CreateOrderPercents")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateOrderPercentsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "goodorderpercent", "crud", "Create")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create goodorderpercent: %v", err)
		return &npool.CreateOrderPercentsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateOrderPercentsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateOrderPercent(
	ctx context.Context,
	in *npool.UpdateOrderPercentRequest,
) (
	*npool.UpdateOrderPercentResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "UpdateOrderPercent")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		return &npool.UpdateOrderPercentResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "goodorderpercent", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create goodorderpercent: %v", err.Error())
		return &npool.UpdateOrderPercentResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateOrderPercentResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetOrderPercent(ctx context.Context, in *npool.GetOrderPercentRequest) (*npool.GetOrderPercentResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetOrderPercent")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.GetOrderPercentResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "goodorderpercent", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get goodorderpercent: %v", err)
		return &npool.GetOrderPercentResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderPercentResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func ValidateConds(conds *npool.Conds) error {
	if conds.ID != nil {
		if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
			return err
		}
	}
	if conds.AppID != nil {
		if _, err := uuid.Parse(conds.GetAppID().GetValue()); err != nil {
			return err
		}
	}
	if conds.UserID != nil {
		if _, err := uuid.Parse(conds.GetUserID().GetValue()); err != nil {
			return err
		}
	}
	if conds.GoodID != nil {
		if _, err := uuid.Parse(conds.GetGoodID().GetValue()); err != nil {
			return err
		}
	}
	for _, id := range conds.GetUserIDs().GetValue() {
		if _, err := uuid.Parse(id); err != nil {
			return err
		}
	}
	for _, id := range conds.GetGoodIDs().GetValue() {
		if _, err := uuid.Parse(id); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) GetOrderPercentOnly(
	ctx context.Context,
	in *npool.GetOrderPercentOnlyRequest,
) (
	*npool.GetOrderPercentOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetOrderPercentOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetOrderPercentOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "goodorderpercent", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get goodorderpercents: %v", err)
		return &npool.GetOrderPercentOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderPercentOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetOrderPercents(ctx context.Context, in *npool.GetOrderPercentsRequest) (*npool.GetOrderPercentsResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetOrderPercents")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetOrderPercentsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "goodorderpercent", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get goodorderpercents: %v", err)
		return &npool.GetOrderPercentsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderPercentsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistOrderPercent(
	ctx context.Context,
	in *npool.ExistOrderPercentRequest,
) (
	*npool.ExistOrderPercentResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "ExistOrderPercent")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.ExistOrderPercentResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "goodorderpercent", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check goodorderpercent: %v", err)
		return &npool.ExistOrderPercentResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistOrderPercentResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistOrderPercentConds(ctx context.Context,
	in *npool.ExistOrderPercentCondsRequest) (*npool.ExistOrderPercentCondsResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "ExistOrderPercentConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.ExistOrderPercentCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "goodorderpercent", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check goodorderpercent: %v", err)
		return &npool.ExistOrderPercentCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistOrderPercentCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountOrderPercents(
	ctx context.Context,
	in *npool.CountOrderPercentsRequest,
) (
	*npool.CountOrderPercentsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CountOrderPercents")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.CountOrderPercentsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "goodorderpercent", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count goodorderpercents: %v", err)
		return &npool.CountOrderPercentsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountOrderPercentsResponse{
		Info: total,
	}, nil
}
