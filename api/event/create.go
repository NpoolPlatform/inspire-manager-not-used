//nolint:nolintlint,dupl
package event

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	converter "github.com/NpoolPlatform/inspire-manager/pkg/converter/event"
	crud "github.com/NpoolPlatform/inspire-manager/pkg/crud/event"
	commontracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer/event"

	constant "github.com/NpoolPlatform/inspire-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/event"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func ValidateCreate(in *npool.EventReq) error { //nolint
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("ValidateCreate", "AppID", in.GetAppID(), "Error", err)
		return err
	}
	switch in.GetEventType() {
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
		if _, err := uuid.Parse(in.GetGoodID()); err != nil {
			logger.Sugar().Errorw("ValidateCreate", "GoodID", in.GetGoodID(), "Error", err)
			return err
		}
	default:
		logger.Sugar().Errorw("ValidateCreate", "EventType", in.GetEventType())
		return fmt.Errorf("eventtype is invalid")
	}
	for _, id := range in.GetCouponIDs() {
		if _, err := uuid.Parse(id); err != nil {
			logger.Sugar().Errorw("ValidateCreate", "CouponIDs", in.GetCouponIDs(), "Error", err)
			return err
		}
	}
	if _, err := decimal.NewFromString(in.GetCredits()); err != nil {
		logger.Sugar().Errorw("ValidateCreate", "Credits", in.GetCredits(), "Error", err)
		return err
	}
	if _, err := decimal.NewFromString(in.GetCreditsPerUSD()); err != nil {
		logger.Sugar().Errorw("ValidateCreate", "CreditsPerUSD", in.GetCreditsPerUSD(), "Error", err)
		return err
	}
	return nil
}

func ValidateCreateMany(in []*npool.EventReq) error {
	for _, info := range in {
		if err := ValidateCreate(info); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) CreateEvent(ctx context.Context, in *npool.CreateEventRequest) (*npool.CreateEventResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateEvent")
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
		return &npool.CreateEventResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	_info := in.GetInfo()

	switch _info.GetEventType() {
	case basetypes.UsedFor_Purchase:
	default:
		_info.GoodID = nil
	}

	span = commontracer.TraceInvoker(span, "event", "crud", "Create")

	info, err := crud.Create(ctx, _info)
	if err != nil {
		logger.Sugar().Errorw("CreateEvent", "Error", err)
		return &npool.CreateEventResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateEventResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateEvents(ctx context.Context, in *npool.CreateEventsRequest) (*npool.CreateEventsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateEvents")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateEventsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	if err := ValidateCreateMany(in.GetInfos()); err != nil {
		return &npool.CreateEventsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "event", "crud", "CreateBulk")

	_infos := in.GetInfos()

	for _, info := range _infos {
		switch info.GetEventType() {
		case basetypes.UsedFor_Purchase:
		default:
			info.GoodID = nil
		}
	}

	rows, err := crud.CreateBulk(ctx, _infos)
	if err != nil {
		logger.Sugar().Errorw("CreateEvent", "Error", err)
		return &npool.CreateEventsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateEventsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}
