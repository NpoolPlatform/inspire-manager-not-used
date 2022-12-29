package discount

import (
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/discount"

	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.CouponDiscount) *npool.Discount {
	if row == nil {
		return nil
	}
	return &npool.Discount{
		ID:              row.ID.String(),
		AppID:           row.AppID.String(),
		Discount:        row.Discount.String(),
		ReleaseByUserID: row.ReleaseByUserID.String(),
		StartAt:         row.StartAt,
		DurationDays:    row.DurationDays,
		Message:         row.Message,
		Name:            row.Name,
		CreatedAt:       row.CreatedAt,
		UpdatedAt:       row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.CouponDiscount) []*npool.Discount {
	infos := []*npool.Discount{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
