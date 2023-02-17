package detail

import (
	"fmt"

	"github.com/shopspring/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/detail"

	"github.com/google/uuid"
)

func validate(info *npool.DetailReq) error { //nolint
	if info.AppID == nil {
		logger.Sugar().Errorw("validate", "AppID", info.AppID)
		return status.Error(codes.InvalidArgument, "AppID is empty")
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", info.AppID, "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	if info.UserID == nil {
		logger.Sugar().Errorw("validate", "UserID", info.UserID)
		return status.Error(codes.InvalidArgument, "UserID is empty")
	}

	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		logger.Sugar().Errorw("validate", "UserID", info.UserID, "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("UserID is invalid: %v", err))
	}

	if info.DirectContributorID != nil {
		if _, err := uuid.Parse(info.GetDirectContributorID()); err != nil {
			logger.Sugar().Errorw("validate", "DirectContributorID", info.DirectContributorID, "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("DirectContributorID is invalid: %v", err))
		}
	}

	if info.GoodID == nil {
		logger.Sugar().Errorw("validate", "GoodID", info.GoodID)
		return status.Error(codes.InvalidArgument, "GoodID is empty")
	}

	if _, err := uuid.Parse(info.GetGoodID()); err != nil {
		logger.Sugar().Errorw("validate", "GoodID", info.GoodID, "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("GoodID is invalid: %v", err))
	}

	if info.OrderID == nil {
		logger.Sugar().Errorw("validate", "OrderID", info.OrderID)
		return status.Error(codes.InvalidArgument, "OrderID is empty")
	}

	if _, err := uuid.Parse(info.GetOrderID()); err != nil {
		logger.Sugar().Errorw("validate", "OrderID", info.OrderID, "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("OrderID is invalid: %v", err))
	}

	if info.PaymentID == nil {
		logger.Sugar().Errorw("validate", "PaymentID", info.PaymentID)
		return status.Error(codes.InvalidArgument, "PaymentID is empty")
	}

	if _, err := uuid.Parse(info.GetPaymentID()); err != nil {
		logger.Sugar().Errorw("validate", "PaymentID", info.PaymentID, "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("PaymentID is invalid: %v", err))
	}

	if info.CoinTypeID == nil {
		logger.Sugar().Errorw("validate", "CoinTypeID", info.CoinTypeID)
		return status.Error(codes.InvalidArgument, "CoinTypeID is empty")
	}

	if _, err := uuid.Parse(info.GetCoinTypeID()); err != nil {
		logger.Sugar().Errorw("validate", "CoinTypeID", info.CoinTypeID, "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("CoinTypeID is invalid: %v", err))
	}

	if info.PaymentCoinTypeID == nil {
		logger.Sugar().Errorw("validate", "PaymentCoinTypeID", info.PaymentCoinTypeID)
		return status.Error(codes.InvalidArgument, "PaymentCoinTypeID is empty")
	}

	if _, err := uuid.Parse(info.GetPaymentCoinTypeID()); err != nil {
		logger.Sugar().Errorw("validate", "PaymentCoinTypeID", info.PaymentCoinTypeID, "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("PaymentCoinTypeID is invalid: %v", err))
	}

	if info.PaymentCoinUSDCurrency != nil {
		currency, err := decimal.NewFromString(info.GetPaymentCoinUSDCurrency())
		if err != nil {
			logger.Sugar().Errorw("validate", "PaymentCoinUSDCurrency", info.GetPaymentCoinUSDCurrency(), "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("PaymentCoinUSDCurrency is invalid: %v", err))
		}
		if currency.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Errorw("validate", "PaymentCoinUSDCurrency", info.GetPaymentCoinUSDCurrency(), "error", "less than 0")
			return status.Error(codes.InvalidArgument, "PaymentCoinUSDCurrency is less than 0")
		}
	}

	if info.Amount != nil {
		amount, err := decimal.NewFromString(info.GetAmount())
		if err != nil {
			logger.Sugar().Errorw("validate", "Amount", info.GetAmount(), "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Amount is invalid: %v", err))
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Errorw("validate", "Amount", info.GetAmount(), "error", "less than 0")
			return status.Error(codes.InvalidArgument, "Amount is less than 0")
		}
	}

	if info.USDAmount != nil {
		amount, err := decimal.NewFromString(info.GetUSDAmount())
		if err != nil {
			logger.Sugar().Errorw("validate", "USDAmount", info.GetUSDAmount(), "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("USDAmount is invalid: %v", err))
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Errorw("validate", "USDAmount", info.GetUSDAmount(), "error", "less than 0")
			return status.Error(codes.InvalidArgument, "USDAmount is less than 0")
		}
	}

	if info.Commission != nil {
		amount, err := decimal.NewFromString(info.GetCommission())
		if err != nil {
			logger.Sugar().Errorw("validate", "Commission", info.GetCommission(), "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Commission is invalid: %v", err))
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Errorw("validate", "Commission", info.GetCommission(), "error", "less than 0")
			return status.Error(codes.InvalidArgument, "Commission is less than 0")
		}
	}

	if info.Units == nil {
		logger.Sugar().Errorw("validate", "Units", info.Units)
		return status.Error(codes.InvalidArgument, "Units is 0")
	}
	units, err := decimal.NewFromString(info.GetUnits())
	if err != nil {
		logger.Sugar().Errorw("validate", "Units", info.Units, "err", err.Error())
		return status.Error(codes.InvalidArgument, err.Error())
	}
	if units.Cmp(decimal.NewFromInt(0)) < 0 {
		logger.Sugar().Errorw("validate", "Units", info.Units)
		return status.Error(codes.InvalidArgument, "Units is 0")
	}
	return nil
}

func duplicate(infos []*npool.DetailReq) error {
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

func Validate(info *npool.DetailReq) error {
	return validate(info)
}
