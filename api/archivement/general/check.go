package general

import (
	"fmt"

	"github.com/shopspring/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/general"

	"github.com/google/uuid"
)

func validate(info *npool.GeneralReq) error { //nolint
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

	if info.GoodID == nil {
		logger.Sugar().Errorw("validate", "GoodID", info.GoodID)
		return status.Error(codes.InvalidArgument, "GoodID is empty")
	}

	if _, err := uuid.Parse(info.GetGoodID()); err != nil {
		logger.Sugar().Errorw("validate", "GoodID", info.GetGoodID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("GoodID is invalid: %v", err))
	}

	if info.CoinTypeID == nil {
		logger.Sugar().Errorw("validate", "CoinTypeID", info.CoinTypeID)
		return status.Error(codes.InvalidArgument, "CoinTypeID is empty")
	}

	if _, err := uuid.Parse(info.GetCoinTypeID()); err != nil {
		logger.Sugar().Errorw("validate", "CoinTypeID", info.GetCoinTypeID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("CoinTypeID is invalid: %v", err))
	}

	if info.TotalAmount != nil {
		amount, err := decimal.NewFromString(info.GetTotalAmount())
		if err != nil {
			logger.Sugar().Errorw("validate", "TotalAmount", info.GetTotalAmount(), "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("TotalAmount is invalid: %v", err))
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Errorw("validate", "TotalAmount", info.GetTotalAmount(), "error", "less than 0")
			return status.Error(codes.InvalidArgument, "TotalAmount is less than 0")
		}
	}

	if info.SelfAmount != nil {
		amount, err := decimal.NewFromString(info.GetSelfAmount())
		if err != nil {
			logger.Sugar().Errorw("validate", "SelfAmount", info.GetSelfAmount(), "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("SelfAmount is invalid: %v", err))
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Errorw("validate", "SelfAmount", info.GetSelfAmount(), "error", "less than 0")
			return status.Error(codes.InvalidArgument, "SelfAmount is less than 0")
		}
	}

	if info.TotalCommission != nil {
		amount, err := decimal.NewFromString(info.GetTotalCommission())
		if err != nil {
			logger.Sugar().Errorw("validate", "TotalCommission", info.GetTotalCommission(), "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("TotalCommission is invalid: %v", err))
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Errorw("validate", "TotalCommission", info.GetTotalCommission(), "error", "less than 0")
			return status.Error(codes.InvalidArgument, "TotalCommission is less than 0")
		}
	}

	if info.SelfCommission != nil {
		amount, err := decimal.NewFromString(info.GetSelfCommission())
		if err != nil {
			logger.Sugar().Errorw("validate", "SelfCommission", info.GetSelfCommission(), "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("SelfCommission is invalid: %v", err))
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Errorw("validate", "SelfCommission", info.GetSelfCommission(), "error", "less than 0")
			return status.Error(codes.InvalidArgument, "SelfCommission is less than 0")
		}
	}

	return nil
}

func duplicate(infos []*npool.GeneralReq) error {
	keys := map[string]struct{}{}
	apps := map[string]struct{}{}

	for _, info := range infos {
		if err := validate(info); err != nil {
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Infos has invalid element %v", err))
		}

		key := fmt.Sprintf("%v:%v:%v", info.AppID, info.UserID, info.CoinTypeID)
		if _, ok := keys[key]; ok {
			return status.Error(codes.InvalidArgument, "Infos has duplicate AppID:UserID:CoinTypeID")
		}

		keys[key] = struct{}{}
		apps[info.GetAppID()] = struct{}{}
	}

	if len(apps) > 1 {
		return status.Error(codes.InvalidArgument, "Infos has different AppID")
	}

	return nil
}
