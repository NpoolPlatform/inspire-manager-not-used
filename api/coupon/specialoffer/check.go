package specialoffer

import (
	"fmt"

	"github.com/shopspring/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/specialoffer"

	"github.com/google/uuid"
)

func validate(info *npool.SpecialOfferReq) error {
	if info.AppID == nil {
		logger.Sugar().Errorw("validate", "AppID", info.AppID)
		return status.Error(codes.InvalidArgument, "AppID is empty")
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", info.GetAppID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	if info.Amount == nil {
		val, err := decimal.NewFromString(info.GetAmount())
		if err != nil {
			logger.Sugar().Errorw("validate", "Amount", info.GetAmount(), "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Amount is invalid: %v", err))
		}
		if val.Cmp(decimal.NewFromInt(0)) <= 0 {
			logger.Sugar().Errorw("validate", "Amount", info.GetAmount(), "error", "Amount less than 0")
			return status.Error(codes.InvalidArgument, "Amount less than 0")
		}
	}
	if info.ReleasedByUserID == nil {
		logger.Sugar().Errorw("validate", "ReleasedByUserID", info.ReleasedByUserID)
		return status.Error(codes.InvalidArgument, "ReleasedByUserID is empty")
	}
	if _, err := uuid.Parse(info.GetReleasedByUserID()); err != nil {
		logger.Sugar().Errorw("validate", "ReleasedByUserID", info.GetReleasedByUserID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("ReleasedByUserID is invalid: %v", err))
	}
	if info.GetStartAt() <= 0 {
		logger.Sugar().Errorw("validate", "Circulation", info.GetStartAt(), "error", "StartAt less than 0")
		return status.Error(codes.InvalidArgument, "Circulation less than 0")
	}
	if info.GetDurationDays() <= 0 {
		logger.Sugar().Errorw("validate", "Circulation", info.GetDurationDays(), "error", "DurationDays less than 0")
		return status.Error(codes.InvalidArgument, "DurationDays less than 0")
	}

	return nil
}
