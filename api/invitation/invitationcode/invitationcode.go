//nolint:dupl
package invitationcode

import (
	"context"
	"fmt"

	converter "github.com/NpoolPlatform/inspire-manager/pkg/converter/invitation/invitationcode"
	crud "github.com/NpoolPlatform/inspire-manager/pkg/crud/invitation/invitationcode"
	commontracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/inspire-manager/pkg/tracer/invitation/invitationcode"

	constant "github.com/NpoolPlatform/inspire-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/inspire/invitation/invitationcode"

	"github.com/google/uuid"
)

func ValidateCreate(info *npool.InvitationCodeReq) error {
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

	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		logger.Sugar().Errorw("validate", "UserID", info.UserID, "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("UserID is invalid: %v", err))
	}

	return nil
}

func (s *Server) CreateInvitationCode(
	ctx context.Context,
	in *npool.CreateInvitationCodeRequest,
) (
	*npool.CreateInvitationCodeResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateInvitationCode")
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
		return &npool.CreateInvitationCodeResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "invitationcode", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create invitationcode: %v", err.Error())
		return &npool.CreateInvitationCodeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateInvitationCodeResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateInvitationCodes(
	ctx context.Context,
	in *npool.CreateInvitationCodesRequest,
) (
	*npool.CreateInvitationCodesResponse,
	error,
) {
	return &npool.CreateInvitationCodesResponse{}, status.Error(codes.Unimplemented, "NOT SUPPORTED")
}

func (s *Server) UpdateInvitationCode(
	ctx context.Context,
	in *npool.UpdateInvitationCodeRequest,
) (
	*npool.UpdateInvitationCodeResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateInvitationCode")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		return &npool.UpdateInvitationCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "invitationcode", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create invitationcode: %v", err.Error())
		return &npool.UpdateInvitationCodeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateInvitationCodeResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetInvitationCode(ctx context.Context, in *npool.GetInvitationCodeRequest) (*npool.GetInvitationCodeResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetInvitationCode")
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
		return &npool.GetInvitationCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "invitationcode", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get invitationcode: %v", err)
		return &npool.GetInvitationCodeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetInvitationCodeResponse{
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
	if conds.InvitationCode != nil && conds.GetInvitationCode().GetValue() == "" {
		return fmt.Errorf("invitationcode is invalid")
	}

	return nil
}

func (s *Server) GetInvitationCodeOnly(
	ctx context.Context,
	in *npool.GetInvitationCodeOnlyRequest,
) (
	*npool.GetInvitationCodeOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetInvitationCodeOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetInvitationCodeOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "invitationcode", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get invitationcodes: %v", err)
		return &npool.GetInvitationCodeOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetInvitationCodeOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetInvitationCodes(ctx context.Context, in *npool.GetInvitationCodesRequest) (*npool.GetInvitationCodesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetInvitationCodes")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetInvitationCodesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "invitationcode", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get invitationcodes: %v", err)
		return &npool.GetInvitationCodesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetInvitationCodesResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistInvitationCode(
	ctx context.Context,
	in *npool.ExistInvitationCodeRequest,
) (
	*npool.ExistInvitationCodeResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistInvitationCode")
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
		return &npool.ExistInvitationCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "invitationcode", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check invitationcode: %v", err)
		return &npool.ExistInvitationCodeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistInvitationCodeResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistInvitationCodeConds(ctx context.Context,
	in *npool.ExistInvitationCodeCondsRequest) (*npool.ExistInvitationCodeCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistInvitationCodeConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.ExistInvitationCodeCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "invitationcode", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check invitationcode: %v", err)
		return &npool.ExistInvitationCodeCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistInvitationCodeCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountInvitationCodes(
	ctx context.Context,
	in *npool.CountInvitationCodesRequest,
) (
	*npool.CountInvitationCodesResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountInvitationCodes")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.CountInvitationCodesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "invitationcode", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count invitationcodes: %v", err)
		return &npool.CountInvitationCodesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountInvitationCodesResponse{
		Info: total,
	}, nil
}
