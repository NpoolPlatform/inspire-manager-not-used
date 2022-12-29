package goodorderpercent

import (
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission/goodorderpercent"
)

func Ent2Grpc(row *ent.GoodOrderPercent) *npool.OrderPercent {
	if row == nil {
		return nil
	}

	return &npool.OrderPercent{
		ID:        row.ID.String(),
		AppID:     row.AppID.String(),
		UserID:    row.UserID.String(),
		GoodID:    row.GoodID.String(),
		Percent:   row.Percent.String(),
		StartAt:   row.StartAt,
		EndAt:     row.EndAt,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.GoodOrderPercent) []*npool.OrderPercent {
	infos := []*npool.OrderPercent{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
