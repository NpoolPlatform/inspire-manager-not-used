//nolint:nolintlint,dupl
package event

import (
	"context"
	"fmt"

	converter "github.com/NpoolPlatform/inspire-manager/pkg/converter/event"
	crud "github.com/NpoolPlatform/inspire-manager/pkg/crud/event"
	commontracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer/event"

	constant "github.com/NpoolPlatform/inspire-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/event"

	"github.com/google/uuid"
)

func ValidateConds(conds *npool.Conds) error { //nolint
	if conds.ID != nil {
		if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
			logger.Sugar().Errorw("ValidateConds", "ID", conds.GetID().GetValue(), "Error", err)
			return err
		}
	}
	if conds.AppID != nil {
		if _, err := uuid.Parse(conds.GetAppID().GetValue()); err != nil {
			logger.Sugar().Errorw("ValidateConds", "AppID", conds.GetAppID().GetValue(), "Error", err)
			return err
		}
	}
	if conds.EventType != nil {
		switch basetypes.UsedFor(conds.GetEventType().GetValue()) {
		case basetypes.UsedFor_Signup:
		case basetypes.UsedFor_Signin:
		case basetypes.UsedFor_Update:
		case basetypes.UsedFor_Contact:
		case basetypes.UsedFor_SetWithdrawAddress:
		case basetypes.UsedFor_Withdraw:
		case basetypes.UsedFor_CreateInvitationCode:
		case basetypes.UsedFor_SetCommission:
		case basetypes.UsedFor_SetTransferTargetUser:
		case basetypes.UsedFor_WithdrawalRequest:
		case basetypes.UsedFor_WithdrawalCompleted:
		case basetypes.UsedFor_DepositReceived:
		case basetypes.UsedFor_KYCApproved:
		case basetypes.UsedFor_KYCRejected:
		case basetypes.UsedFor_Purchase:
		case basetypes.UsedFor_AffiliatePurchase:
		default:
			logger.Sugar().Errorw("ValidateConds", "EventType", conds.GetEventType().GetValue())
			return fmt.Errorf("eventtype is invalid")
		}
	}
	return nil
}

func (s *Server) GetEvent(ctx context.Context, in *npool.GetEventRequest) (*npool.GetEventResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetEvent")
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
		logger.Sugar().Errorw("GetEvent", "ID", in.GetID(), "Error", err)
		return &npool.GetEventResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "event", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetEvent", "Error", err)
		return &npool.GetEventResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetEventResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetEventOnly(ctx context.Context, in *npool.GetEventOnlyRequest) (*npool.GetEventOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetEventOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetEventOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "event", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetEventOnly", "Error", err)
		return &npool.GetEventOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetEventOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetEvents(ctx context.Context, in *npool.GetEventsRequest) (*npool.GetEventsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetEvents")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetEventsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "event", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetEvents", "Error", err)
		return &npool.GetEventsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetEventsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistEvent(ctx context.Context, in *npool.ExistEventRequest) (*npool.ExistEventResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistEvent")
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
		logger.Sugar().Errorw("ExistEvent", "ID", in.GetID(), "Error", err)
		return &npool.ExistEventResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "event", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistEvent", "Error", err)
		return &npool.ExistEventResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistEventResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistEventConds(ctx context.Context,
	in *npool.ExistEventCondsRequest) (*npool.ExistEventCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistEventConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.ExistEventCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "event", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistEventConds", "Error", err)
		return &npool.ExistEventCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistEventCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountEvents(ctx context.Context, in *npool.CountEventsRequest) (*npool.CountEventsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountEvents")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.CountEventsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "event", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("CountEvents", "Error", err)
		return &npool.CountEventsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountEventsResponse{
		Info: total,
	}, nil
}
