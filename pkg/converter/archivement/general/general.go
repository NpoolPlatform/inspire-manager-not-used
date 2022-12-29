package general

import (
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/general"
)

func Ent2Grpc(row *ent.ArchivementGeneral) *npool.General {
	if row == nil {
		return nil
	}

	return &npool.General{
		ID:              row.ID.String(),
		AppID:           row.AppID.String(),
		UserID:          row.UserID.String(),
		GoodID:          row.GoodID.String(),
		CoinTypeID:      row.CoinTypeID.String(),
		TotalAmount:     row.TotalAmount.String(),
		SelfAmount:      row.SelfAmount.String(),
		TotalUnits:      row.TotalUnits,
		SelfUnits:       row.SelfUnits,
		TotalCommission: row.TotalCommission.String(),
		SelfCommission:  row.SelfCommission.String(),
	}
}

func Ent2GrpcMany(rows []*ent.ArchivementGeneral) []*npool.General {
	infos := []*npool.General{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
