//nolint:dupl
package goodordervaluepercent

import (
	"context"
	"fmt"

	converter "github.com/NpoolPlatform/inspire-manager/pkg/converter/commission/goodordervaluepercent"
	crud "github.com/NpoolPlatform/inspire-manager/pkg/crud/commission/goodordervaluepercent"
	commontracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer/commission/goodordervaluepercent"

	constant "github.com/NpoolPlatform/inspire-manager/pkg/message/const"

	"github.com/shopspring/decimal"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission/goodordervaluepercent"

	"github.com/google/uuid"
)

func ValidateCreate(info *npool.OrderValuePercentReq) error {
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

func (s *Server) CreateOrderValuePercent(
	ctx context.Context,
	in *npool.CreateOrderValuePercentRequest,
) (
	*npool.CreateOrderValuePercentResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateOrderValuePercent")
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
		return &npool.CreateOrderValuePercentResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "goodordervaluepercent", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create goodordervaluepercent: %v", err.Error())
		return &npool.CreateOrderValuePercentResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateOrderValuePercentResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateOrderValuePercents(
	ctx context.Context,
	in *npool.CreateOrderValuePercentsRequest,
) (
	*npool.CreateOrderValuePercentsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateOrderValuePercents")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateOrderValuePercentsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "goodordervaluepercent", "crud", "Create")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create goodordervaluepercent: %v", err)
		return &npool.CreateOrderValuePercentsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateOrderValuePercentsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateOrderValuePercent(
	ctx context.Context,
	in *npool.UpdateOrderValuePercentRequest,
) (
	*npool.UpdateOrderValuePercentResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateOrderValuePercent")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		return &npool.UpdateOrderValuePercentResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "goodordervaluepercent", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create goodordervaluepercent: %v", err.Error())
		return &npool.UpdateOrderValuePercentResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateOrderValuePercentResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetOrderValuePercent(
	ctx context.Context,
	in *npool.GetOrderValuePercentRequest,
) (*npool.GetOrderValuePercentResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetOrderValuePercent")
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
		return &npool.GetOrderValuePercentResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "goodordervaluepercent", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get goodordervaluepercent: %v", err)
		return &npool.GetOrderValuePercentResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderValuePercentResponse{
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

func (s *Server) GetOrderValuePercentOnly(
	ctx context.Context,
	in *npool.GetOrderValuePercentOnlyRequest,
) (
	*npool.GetOrderValuePercentOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetOrderValuePercentOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetOrderValuePercentOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "goodordervaluepercent", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get goodordervaluepercents: %v", err)
		return &npool.GetOrderValuePercentOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderValuePercentOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetOrderValuePercents(
	ctx context.Context,
	in *npool.GetOrderValuePercentsRequest,
) (*npool.GetOrderValuePercentsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetOrderValuePercents")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetOrderValuePercentsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "goodordervaluepercent", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get goodordervaluepercents: %v", err)
		return &npool.GetOrderValuePercentsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderValuePercentsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistOrderValuePercent(
	ctx context.Context,
	in *npool.ExistOrderValuePercentRequest,
) (
	*npool.ExistOrderValuePercentResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistOrderValuePercent")
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
		return &npool.ExistOrderValuePercentResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "goodordervaluepercent", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check goodordervaluepercent: %v", err)
		return &npool.ExistOrderValuePercentResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistOrderValuePercentResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistOrderValuePercentConds(ctx context.Context,
	in *npool.ExistOrderValuePercentCondsRequest) (*npool.ExistOrderValuePercentCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistOrderValuePercentConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.ExistOrderValuePercentCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "goodordervaluepercent", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check goodordervaluepercent: %v", err)
		return &npool.ExistOrderValuePercentCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistOrderValuePercentCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountOrderValuePercents(
	ctx context.Context,
	in *npool.CountOrderValuePercentsRequest,
) (
	*npool.CountOrderValuePercentsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountOrderValuePercents")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.CountOrderValuePercentsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "goodordervaluepercent", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count goodordervaluepercents: %v", err)
		return &npool.CountOrderValuePercentsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountOrderValuePercentsResponse{
		Info: total,
	}, nil
}
