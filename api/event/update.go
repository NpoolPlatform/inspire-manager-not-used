//nolint:nolintlint,dupl
package event

import (
	"context"

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
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/event"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func ValidateUpdate(in *npool.EventReq) error {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("ValidateUpdate", "ID", in.GetID(), "Error", err)
		return err
	}
	for _, id := range in.GetCouponIDs() {
		if _, err := uuid.Parse(id); err != nil {
			logger.Sugar().Errorw("ValidateUpdate", "CouponIDs", in.GetCouponIDs(), "Error", err)
			return err
		}
	}
	if in.Credits != nil {
		if _, err := decimal.NewFromString(in.GetCredits()); err != nil {
			logger.Sugar().Errorw("ValidateUpdate", "Credits", in.GetCredits(), "Error", err)
			return err
		}
	}
	if in.CreditsPerUSD != nil {
		if _, err := decimal.NewFromString(in.GetCreditsPerUSD()); err != nil {
			logger.Sugar().Errorw("ValidateUpdate", "CreditsPerUSD", in.GetCreditsPerUSD(), "Error", err)
			return err
		}
	}
	return nil
}

func (s *Server) UpdateEvent(ctx context.Context, in *npool.UpdateEventRequest) (*npool.UpdateEventResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateEvent")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if err := ValidateUpdate(in.GetInfo()); err != nil {
		return &npool.UpdateEventResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "event", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateEvent", "Error", err)
		return &npool.UpdateEventResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateEventResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
