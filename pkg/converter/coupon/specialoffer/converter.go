package specialoffer

import (
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/specialoffer"

	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.CouponSpecialOffer) *npool.SpecialOffer {
	if row == nil {
		return nil
	}
	return &npool.SpecialOffer{
		ID:              row.ID.String(),
		AppID:           row.AppID.String(),
		UserID:          row.UserID.String(),
		Amount:          row.Amount.String(),
		ReleaseByUserID: row.ReleaseByUserID.String(),
		StartAt:         row.StartAt,
		DurationDays:    row.DurationDays,
		Message:         row.Message,
		CreatedAt:       row.CreatedAt,
		UpdatedAt:       row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.CouponSpecialOffer) []*npool.SpecialOffer {
	infos := []*npool.SpecialOffer{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
