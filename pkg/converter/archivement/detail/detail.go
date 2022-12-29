package detail

import (
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/detail"
)

func Ent2Grpc(row *ent.ArchivementDetail) *npool.Detail {
	if row == nil {
		return nil
	}

	return &npool.Detail{
		ID:                     row.ID.String(),
		AppID:                  row.AppID.String(),
		UserID:                 row.UserID.String(),
		DirectContributorID:    row.DirectContributorID.String(),
		GoodID:                 row.GoodID.String(),
		OrderID:                row.OrderID.String(),
		SelfOrder:              row.SelfOrder,
		PaymentID:              row.PaymentID.String(),
		CoinTypeID:             row.CoinTypeID.String(),
		PaymentCoinTypeID:      row.PaymentCoinTypeID.String(),
		PaymentCoinUSDCurrency: row.PaymentCoinUsdCurrency.String(),
		Units:                  row.Units,
		Amount:                 row.Amount.String(),
		USDAmount:              row.UsdAmount.String(),
		Commission:             row.Commission.String(),
		CreatedAt:              row.CreatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.ArchivementDetail) []*npool.Detail {
	infos := []*npool.Detail{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
