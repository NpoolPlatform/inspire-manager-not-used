package allocated

import (
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"

	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.CouponAllocated) *npool.Allocated {
	if row == nil {
		return nil
	}
	return &npool.Allocated{
		ID:        row.ID.String(),
		AppID:     row.AppID.String(),
		UserID:    row.UserID.String(),
		Type:      npool.CouponType(npool.CouponType_value[row.Type]),
		CouponID:  row.CouponID.String(),
		Value:     row.Value.String(),
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.CouponAllocated) []*npool.Allocated {
	infos := []*npool.Allocated{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
