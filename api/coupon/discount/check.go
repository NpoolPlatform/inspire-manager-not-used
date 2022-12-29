package discount

import (
	"fmt"

	"github.com/shopspring/decimal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/discount"

	"github.com/google/uuid"
)

func validate(info *npool.DiscountReq) error {
	if info.AppID == nil {
		logger.Sugar().Errorw("validate", "AppID", info.AppID)
		return status.Error(codes.InvalidArgument, "AppID is empty")
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", info.GetAppID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	if info.Discount == nil {
		val, err := decimal.NewFromString(info.GetDiscount())
		if err != nil {
			logger.Sugar().Errorw("validate", "Discount", info.GetDiscount(), "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Discount is invalid: %v", err))
		}
		if val.Cmp(decimal.NewFromInt(0)) <= 0 {
			logger.Sugar().Errorw("validate", "Discount", info.GetDiscount(), "error", "Discount less than 0")
			return status.Error(codes.InvalidArgument, "Discount less than 0")
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
		logger.Sugar().Errorw("validate", "Circulation", info.GetStartAt(), "error", "StartAt less than 0")
		return status.Error(codes.InvalidArgument, "Circulation less than 0")
	}
	if info.GetDurationDays() <= 0 {
		logger.Sugar().Errorw("validate", "Circulation", info.GetDurationDays(), "error", "DurationDays less than 0")
		return status.Error(codes.InvalidArgument, "DurationDays less than 0")
	}

	if info.GetName() == "" {
		logger.Sugar().Errorw("validate", "Circulation", info.GetName(), "error", "Name is empty")
		return status.Error(codes.InvalidArgument, "Name is empty")
	}

	return nil
}
