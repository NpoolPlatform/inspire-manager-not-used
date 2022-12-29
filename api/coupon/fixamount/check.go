package fixamount

import (
	"fmt"

	"github.com/shopspring/decimal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/fixamount"

	"github.com/google/uuid"
)

func validate(info *npool.FixAmountReq) error {
	if info.AppID == nil {
		logger.Sugar().Errorw("validate", "AppID", info.AppID)
		return status.Error(codes.InvalidArgument, "AppID is empty")
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", info.GetAppID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	if info.Denomination == nil {
		val, err := decimal.NewFromString(info.GetDenomination())
		if err != nil {
			logger.Sugar().Errorw("validate", "Denomination", info.GetDenomination(), "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Denomination is invalid: %v", err))
		}
		if val.Cmp(decimal.NewFromInt(0)) <= 0 {
			logger.Sugar().Errorw("validate", "Denomination", info.GetDenomination(), "error", "Denomination less than 0")
			return status.Error(codes.InvalidArgument, "Denomination less than 0")
		}
	}
	if info.Circulation == nil {
		val, err := decimal.NewFromString(info.GetCirculation())
		if err != nil {
			logger.Sugar().Errorw("validate", "Circulation", info.GetCirculation(), "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Circulation is invalid: %v", err))
		}
		if val.Cmp(decimal.NewFromInt(0)) <= 0 {
			logger.Sugar().Errorw("validate", "Circulation", info.GetCirculation(), "error", "Circulation less than 0")
			return status.Error(codes.InvalidArgument, "Circulation less than 0")
		}
	}
	if info.ReleaseByUserID == nil {
		logger.Sugar().Errorw("validate", "ReleaseByUserID", info.ReleaseByUserID)
		return status.Error(codes.InvalidArgument, "ReleaseByUserID is empty")
	}
	if _, err := uuid.Parse(info.GetReleaseByUserID()); err != nil {
		logger.Sugar().Errorw("validate", "ReleaseByUserID", info.GetReleaseByUserID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("ReleaseByUserID is invalid: %v", err))
	}
	if info.GetStartAt() <= 0 {
		logger.Sugar().Errorw("validate", "Circulation", info.GetCirculation(), "error", "Circulation less than 0")
		return status.Error(codes.InvalidArgument, "Circulation less than 0")
	}
	if info.GetDurationDays() <= 0 {
		logger.Sugar().Errorw("validate", "Circulation", info.GetDurationDays(), "error", "DurationDays less than 0")
		return status.Error(codes.InvalidArgument, "DurationDays less than 0")
	}

	if info.GetName() == "" {
		logger.Sugar().Errorw("validate", "Circulation", info.GetName(), "error", "NameName less than 0")
		return status.Error(codes.InvalidArgument, "Name is empty")
	}

	return nil
}
