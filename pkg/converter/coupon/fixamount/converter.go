package fixamount

import (
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/inspire/coupon/fixamount"

	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.CouponFixAmount) *npool.FixAmount {
	if row == nil {
		return nil
	}
	return &npool.FixAmount{
		ID:              row.ID.String(),
		AppID:           row.AppID.String(),
		Denomination:    row.Denomination.String(),
		Circulation:     row.Circulation.String(),
		ReleaseByUserID: row.ReleaseByUserID.String(),
		StartAt:         row.StartAt,
		DurationDays:    row.DurationDays,
		Message:         row.Message,
		Name:            row.Name,
		CreatedAt:       row.CreatedAt,
		UpdatedAt:       row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.CouponFixAmount) []*npool.FixAmount {
	infos := []*npool.FixAmount{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
