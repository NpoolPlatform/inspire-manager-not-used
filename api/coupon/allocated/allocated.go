//nolint:nolintlint,dupl
package allocated

import (
	"context"

	converter "github.com/NpoolPlatform/inspire-manager/pkg/converter/coupon/allocated"
	crud "github.com/NpoolPlatform/inspire-manager/pkg/crud/coupon/allocated"
	commontracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer/coupon/allocated"

	constant "github.com/NpoolPlatform/inspire-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"

	"github.com/google/uuid"
)

func (s *Server) CreateAllocated(ctx context.Context, in *npool.CreateAllocatedRequest) (*npool.CreateAllocatedResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAllocated")
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
		return &npool.CreateAllocatedResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "allocated", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create allocated: %v", err.Error())
		return &npool.CreateAllocatedResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAllocatedResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateAllocateds(ctx context.Context, in *npool.CreateAllocatedsRequest) (*npool.CreateAllocatedsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAllocateds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateAllocatedsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "allocated", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create allocateds: %v", err)
		return &npool.CreateAllocatedsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAllocatedsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateAllocated(ctx context.Context, in *npool.UpdateAllocatedRequest) (*npool.UpdateAllocatedResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateAllocated")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		return &npool.UpdateAllocatedResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.GetInfo().UsedByOrderID != nil {
		if _, err := uuid.Parse(in.GetInfo().GetUsedByOrderID()); err != nil {
			return &npool.UpdateAllocatedResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	span = commontracer.TraceInvoker(span, "allocated", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create allocated: %v", err.Error())
		return &npool.UpdateAllocatedResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAllocatedResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAllocated(ctx context.Context, in *npool.GetAllocatedRequest) (*npool.GetAllocatedResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAllocated")
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
		return &npool.GetAllocatedResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "allocated", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get allocated: %v", err)
		return &npool.GetAllocatedResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAllocatedResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAllocatedOnly(ctx context.Context, in *npool.GetAllocatedOnlyRequest) (*npool.GetAllocatedOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAllocatedOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "allocated", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get allocateds: %v", err)
		return &npool.GetAllocatedOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAllocatedOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAllocateds(ctx context.Context, in *npool.GetAllocatedsRequest) (*npool.GetAllocatedsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAllocateds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "allocated", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get allocateds: %v", err)
		return &npool.GetAllocatedsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAllocatedsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistAllocated(ctx context.Context, in *npool.ExistAllocatedRequest) (*npool.ExistAllocatedResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAllocated")
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
		return &npool.ExistAllocatedResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "allocated", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check allocated: %v", err)
		return &npool.ExistAllocatedResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAllocatedResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistAllocatedConds(ctx context.Context,
	in *npool.ExistAllocatedCondsRequest) (*npool.ExistAllocatedCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAllocatedConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "allocated", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check allocated: %v", err)
		return &npool.ExistAllocatedCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAllocatedCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountAllocateds(ctx context.Context, in *npool.CountAllocatedsRequest) (*npool.CountAllocatedsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountAllocateds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "allocated", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count allocateds: %v", err)
		return &npool.CountAllocatedsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAllocatedsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteAllocated(ctx context.Context, in *npool.DeleteAllocatedRequest) (*npool.DeleteAllocatedResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAllocateds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "allocated", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.DeleteAllocatedResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Delete(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorf("fail create allocateds: %v", err)
		return &npool.DeleteAllocatedResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAllocatedResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}
