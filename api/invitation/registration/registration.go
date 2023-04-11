//nolint:dupl
package registration

import (
	"context"
	"fmt"

	converter "github.com/NpoolPlatform/inspire-manager/pkg/converter/invitation/registration"
	crud "github.com/NpoolPlatform/inspire-manager/pkg/crud/invitation/registration"
	commontracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer/invitation/registration"

	"github.com/NpoolPlatform/inspire-manager/pkg/servicename"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"

	"github.com/google/uuid"
)

func ValidateCreate(info *npool.RegistrationReq) error {
	if info.ID != nil {
		if _, err := uuid.Parse(info.GetID()); err != nil {
			logger.Sugar().Errorw("validate", "ID", info.ID, "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("ID is invalid: %v", err))
		}
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", info.AppID, "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	if _, err := uuid.Parse(info.GetInviterID()); err != nil {
		logger.Sugar().Errorw("validate", "InviterID", info.InviterID, "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("InviterID is invalid: %v", err))
	}
	if _, err := uuid.Parse(info.GetInviteeID()); err != nil {
		logger.Sugar().Errorw("validate", "InviteeID", info.InviteeID, "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("InviteeID is invalid: %v", err))
	}

	return nil
}

func (s *Server) CreateRegistration(
	ctx context.Context,
	in *npool.CreateRegistrationRequest,
) (
	*npool.CreateRegistrationResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CreateRegistration")
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
		return &npool.CreateRegistrationResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "registration", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create registration: %v", err.Error())
		return &npool.CreateRegistrationResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateRegistrationResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateRegistrations(
	ctx context.Context,
	in *npool.CreateRegistrationsRequest,
) (
	*npool.CreateRegistrationsResponse,
	error,
) {
	return &npool.CreateRegistrationsResponse{}, status.Error(codes.Unimplemented, "NOT SUPPORTED")
}

func (s *Server) UpdateRegistration(
	ctx context.Context,
	in *npool.UpdateRegistrationRequest,
) (
	*npool.UpdateRegistrationResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "UpdateRegistration")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		return &npool.UpdateRegistrationResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "registration", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create registration: %v", err.Error())
		return &npool.UpdateRegistrationResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateRegistrationResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetRegistration(ctx context.Context, in *npool.GetRegistrationRequest) (*npool.GetRegistrationResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetRegistration")
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
		return &npool.GetRegistrationResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "registration", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get registration: %v", err)
		return &npool.GetRegistrationResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetRegistrationResponse{
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
	if conds.InviterID != nil {
		if _, err := uuid.Parse(conds.GetInviterID().GetValue()); err != nil {
			return err
		}
	}
	if conds.InviteeID != nil {
		if _, err := uuid.Parse(conds.GetInviteeID().GetValue()); err != nil {
			return err
		}
	}
	for _, id := range conds.GetInviterIDs().GetValue() {
		if _, err := uuid.Parse(id); err != nil {
			return err
		}
	}
	for _, id := range conds.GetInviteeIDs().GetValue() {
		if _, err := uuid.Parse(id); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) GetRegistrationOnly(
	ctx context.Context,
	in *npool.GetRegistrationOnlyRequest,
) (
	*npool.GetRegistrationOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetRegistrationOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetRegistrationOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "registration", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get registrations: %v", err)
		return &npool.GetRegistrationOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetRegistrationOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetRegistrations(ctx context.Context, in *npool.GetRegistrationsRequest) (*npool.GetRegistrationsResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetRegistrations")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetRegistrationsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "registration", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get registrations: %v", err)
		return &npool.GetRegistrationsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetRegistrationsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistRegistration(
	ctx context.Context,
	in *npool.ExistRegistrationRequest,
) (
	*npool.ExistRegistrationResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "ExistRegistration")
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
		return &npool.ExistRegistrationResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "registration", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check registration: %v", err)
		return &npool.ExistRegistrationResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistRegistrationResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistRegistrationConds(ctx context.Context,
	in *npool.ExistRegistrationCondsRequest) (*npool.ExistRegistrationCondsResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "ExistRegistrationConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.ExistRegistrationCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "registration", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check registration: %v", err)
		return &npool.ExistRegistrationCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistRegistrationCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountRegistrations(
	ctx context.Context,
	in *npool.CountRegistrationsRequest,
) (
	*npool.CountRegistrationsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CountRegistrations")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.CountRegistrationsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "registration", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count registrations: %v", err)
		return &npool.CountRegistrationsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountRegistrationsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteRegistration(
	ctx context.Context,
	in *npool.DeleteRegistrationRequest,
) (
	*npool.DeleteRegistrationResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "DeleteRegistration")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		return &npool.DeleteRegistrationResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "registration", "crud", "Delete")

	info, err := crud.Delete(ctx, uuid.MustParse(in.GetID()))
	if err != nil {
		logger.Sugar().Errorf("fail count registrations: %v", err)
		return &npool.DeleteRegistrationResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteRegistrationResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
