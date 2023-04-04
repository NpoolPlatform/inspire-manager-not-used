//nolint:nolintlint,dupl
package discount

import (
	"context"

	converter "github.com/NpoolPlatform/inspire-manager/pkg/converter/coupon/discount"
	crud "github.com/NpoolPlatform/inspire-manager/pkg/crud/coupon/discount"
	commontracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer/coupon/discount"

	"github.com/NpoolPlatform/inspire-manager/pkg/servicename"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/discount"

	"github.com/google/uuid"
)

func (s *Server) CreateDiscount(ctx context.Context, in *npool.CreateDiscountRequest) (*npool.CreateDiscountResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CreateDiscount")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	err = validate(in.GetInfo())
	if err != nil {
		return &npool.CreateDiscountResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "fixamount", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create fixamount: %v", err.Error())
		return &npool.CreateDiscountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateDiscountResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateDiscounts(ctx context.Context, in *npool.CreateDiscountsRequest) (*npool.CreateDiscountsResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CreateDiscounts")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateDiscountsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "fixamount", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create fixamounts: %v", err)
		return &npool.CreateDiscountsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateDiscountsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateDiscount(
	ctx context.Context,
	in *npool.UpdateDiscountRequest,
) (
	*npool.UpdateDiscountResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "UpdateDiscounts")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())
	span = commontracer.TraceInvoker(span, "fixamount", "crud", "CreateBulk")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create fixamounts: %v", err)
		return &npool.UpdateDiscountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateDiscountResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetDiscount(ctx context.Context, in *npool.GetDiscountRequest) (*npool.GetDiscountResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetDiscount")
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
		return &npool.GetDiscountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fixamount", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get fixamount: %v", err)
		return &npool.GetDiscountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetDiscountResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetDiscountOnly(ctx context.Context, in *npool.GetDiscountOnlyRequest) (*npool.GetDiscountOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetDiscountOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "fixamount", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get fixamounts: %v", err)
		return &npool.GetDiscountOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetDiscountOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetDiscounts(ctx context.Context, in *npool.GetDiscountsRequest) (*npool.GetDiscountsResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetDiscounts")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "fixamount", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get fixamounts: %v", err)
		return &npool.GetDiscountsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetDiscountsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistDiscount(ctx context.Context, in *npool.ExistDiscountRequest) (*npool.ExistDiscountResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "ExistDiscount")
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
		return &npool.ExistDiscountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fixamount", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check fixamount: %v", err)
		return &npool.ExistDiscountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistDiscountResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistDiscountConds(ctx context.Context,
	in *npool.ExistDiscountCondsRequest) (*npool.ExistDiscountCondsResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "ExistDiscountConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "fixamount", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check fixamount: %v", err)
		return &npool.ExistDiscountCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistDiscountCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountDiscounts(ctx context.Context, in *npool.CountDiscountsRequest) (*npool.CountDiscountsResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CountDiscounts")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "fixamount", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count fixamounts: %v", err)
		return &npool.CountDiscountsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountDiscountsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteDiscount(ctx context.Context, in *npool.DeleteDiscountRequest) (*npool.DeleteDiscountResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CreateDiscounts")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "fixamount", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.DeleteDiscountResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Delete(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorf("fail create fixamounts: %v", err)
		return &npool.DeleteDiscountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteDiscountResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}
