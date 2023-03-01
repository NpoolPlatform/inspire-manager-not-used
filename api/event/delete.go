//nolint:nolintlint,dupl
package event

import (
	"context"

	converter "github.com/NpoolPlatform/inspire-manager/pkg/converter/event"
	crud "github.com/NpoolPlatform/inspire-manager/pkg/crud/event"
	commontracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer"

	constant "github.com/NpoolPlatform/inspire-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/event"

	"github.com/google/uuid"
)

func (s *Server) DeleteEvent(ctx context.Context, in *npool.DeleteEventRequest) (*npool.DeleteEventResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteEvent")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.DeleteEventResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	span = commontracer.TraceInvoker(span, "event", "crud", "Delete")

	rows, err := crud.Delete(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteEvent", "Error", err)
		return &npool.DeleteEventResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteEventResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}
