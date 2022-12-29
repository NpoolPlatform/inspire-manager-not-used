package allocated

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"

	"github.com/google/uuid"
)

func validate(info *npool.AllocatedReq) error {
	if info.AppID == nil {
		logger.Sugar().Errorw("validate", "AppID", info.AppID)
		return status.Error(codes.InvalidArgument, "AppID is empty")
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", info.GetAppID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}
	if info.UserID == nil {
		logger.Sugar().Errorw("validate", "UserID", info.UserID)
		return status.Error(codes.InvalidArgument, "UserID is empty")
	}

	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		logger.Sugar().Errorw("validate", "UserID", info.GetUserID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("UserID is invalid: %v", err))
	}
	if info.CouponID == nil {
		logger.Sugar().Errorw("validate", "CouponID", info.CouponID)
		return status.Error(codes.InvalidArgument, "CouponID is empty")
	}

	if _, err := uuid.Parse(info.GetCouponID()); err != nil {
		logger.Sugar().Errorw("validate", "CouponID", info.GetCouponID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("CouponID is invalid: %v", err))
	}
	return nil
}
