//nolint:nolintlint,dupl
package fixamount

import (
	"context"

	converter "github.com/NpoolPlatform/inspire-manager/pkg/converter/coupon/fixamount"
	crud "github.com/NpoolPlatform/inspire-manager/pkg/crud/coupon/fixamount"
	commontracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer/coupon/fixamount"

	"github.com/NpoolPlatform/inspire-manager/pkg/servicename"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/fixamount"

	"github.com/google/uuid"
)

func (s *Server) CreateFixAmount(ctx context.Context, in *npool.CreateFixAmountRequest) (*npool.CreateFixAmountResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CreateFixAmount")
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
		return &npool.CreateFixAmountResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "fixamount", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create fixamount: %v", err.Error())
		return &npool.CreateFixAmountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateFixAmountResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateFixAmounts(ctx context.Context, in *npool.CreateFixAmountsRequest) (*npool.CreateFixAmountsResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CreateFixAmounts")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateFixAmountsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "fixamount", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create fixamounts: %v", err)
		return &npool.CreateFixAmountsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateFixAmountsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateFixAmount(
	ctx context.Context,
	in *npool.UpdateFixAmountRequest,
) (
	*npool.UpdateFixAmountResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "UpdateFixAmounts")
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
		return &npool.UpdateFixAmountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateFixAmountResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetFixAmount(ctx context.Context, in *npool.GetFixAmountRequest) (*npool.GetFixAmountResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetFixAmount")
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
		return &npool.GetFixAmountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fixamount", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get fixamount: %v", err)
		return &npool.GetFixAmountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFixAmountResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetFixAmountOnly(ctx context.Context, in *npool.GetFixAmountOnlyRequest) (*npool.GetFixAmountOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetFixAmountOnly")
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
		return &npool.GetFixAmountOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFixAmountOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetFixAmounts(ctx context.Context, in *npool.GetFixAmountsRequest) (*npool.GetFixAmountsResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetFixAmounts")
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
		return &npool.GetFixAmountsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFixAmountsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistFixAmount(ctx context.Context, in *npool.ExistFixAmountRequest) (*npool.ExistFixAmountResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "ExistFixAmount")
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
		return &npool.ExistFixAmountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fixamount", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check fixamount: %v", err)
		return &npool.ExistFixAmountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistFixAmountResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistFixAmountConds(ctx context.Context,
	in *npool.ExistFixAmountCondsRequest) (*npool.ExistFixAmountCondsResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "ExistFixAmountConds")
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
		return &npool.ExistFixAmountCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistFixAmountCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountFixAmounts(ctx context.Context, in *npool.CountFixAmountsRequest) (*npool.CountFixAmountsResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CountFixAmounts")
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
		return &npool.CountFixAmountsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountFixAmountsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteFixAmount(ctx context.Context, in *npool.DeleteFixAmountRequest) (*npool.DeleteFixAmountResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CreateFixAmounts")
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
		return &npool.DeleteFixAmountResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Delete(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorf("fail create fixamounts: %v", err)
		return &npool.DeleteFixAmountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteFixAmountResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}
