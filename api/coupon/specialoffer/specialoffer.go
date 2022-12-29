//nolint:nolintlint,dupl
package specialoffer

import (
	"context"

	converter "github.com/NpoolPlatform/inspire-manager/pkg/converter/coupon/specialoffer"
	crud "github.com/NpoolPlatform/inspire-manager/pkg/crud/coupon/specialoffer"
	commontracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer/coupon/specialoffer"

	constant "github.com/NpoolPlatform/inspire-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/specialoffer"

	"github.com/google/uuid"
)

func (s *Server) CreateSpecialOffer(
	ctx context.Context,
	in *npool.CreateSpecialOfferRequest,
) (
	*npool.CreateSpecialOfferResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateSpecialOffer")
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
		return &npool.CreateSpecialOfferResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "fixamount", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create fixamount: %v", err.Error())
		return &npool.CreateSpecialOfferResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateSpecialOfferResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateSpecialOffers(
	ctx context.Context,
	in *npool.CreateSpecialOffersRequest,
) (
	*npool.CreateSpecialOffersResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateSpecialOffers")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateSpecialOffersResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "fixamount", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create fixamounts: %v", err)
		return &npool.CreateSpecialOffersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateSpecialOffersResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateSpecialOffer(
	ctx context.Context,
	in *npool.UpdateSpecialOfferRequest,
) (
	*npool.UpdateSpecialOfferResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateSpecialOffers")
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
		return &npool.UpdateSpecialOfferResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateSpecialOfferResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetSpecialOffer(
	ctx context.Context,
	in *npool.GetSpecialOfferRequest,
) (
	*npool.GetSpecialOfferResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetSpecialOffer")
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
		return &npool.GetSpecialOfferResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fixamount", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get fixamount: %v", err)
		return &npool.GetSpecialOfferResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSpecialOfferResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetSpecialOfferOnly(
	ctx context.Context,
	in *npool.GetSpecialOfferOnlyRequest,
) (
	*npool.GetSpecialOfferOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetSpecialOfferOnly")
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
		return &npool.GetSpecialOfferOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSpecialOfferOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetSpecialOffers(ctx context.Context, in *npool.GetSpecialOffersRequest) (*npool.GetSpecialOffersResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetSpecialOffers")
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
		return &npool.GetSpecialOffersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSpecialOffersResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistSpecialOffer(
	ctx context.Context,
	in *npool.ExistSpecialOfferRequest,
) (
	*npool.ExistSpecialOfferResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistSpecialOffer")
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
		return &npool.ExistSpecialOfferResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "fixamount", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check fixamount: %v", err)
		return &npool.ExistSpecialOfferResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistSpecialOfferResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistSpecialOfferConds(
	ctx context.Context,
	in *npool.ExistSpecialOfferCondsRequest,
) (
	*npool.ExistSpecialOfferCondsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistSpecialOfferConds")
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
		return &npool.ExistSpecialOfferCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistSpecialOfferCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountSpecialOffers(
	ctx context.Context,
	in *npool.CountSpecialOffersRequest,
) (
	*npool.CountSpecialOffersResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountSpecialOffers")
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
		return &npool.CountSpecialOffersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountSpecialOffersResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteSpecialOffer(
	ctx context.Context,
	in *npool.DeleteSpecialOfferRequest,
) (
	*npool.DeleteSpecialOfferResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateSpecialOffers")
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
		return &npool.DeleteSpecialOfferResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Delete(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorf("fail create fixamounts: %v", err)
		return &npool.DeleteSpecialOfferResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteSpecialOfferResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}
