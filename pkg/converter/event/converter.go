package event

import (
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/event"

	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.Event) *npool.Event {
	if row == nil {
		return nil
	}

	coupons := []*npool.Coupon{}
	for i, _ := range row.Coupons { //nolint
		coupons = append(coupons, &row.Coupons[i])
	}

	ret := &npool.Event{
		ID:             row.ID.String(),
		AppID:          row.AppID.String(),
		EventType:      basetypes.UsedFor(basetypes.UsedFor_value[row.EventType]),
		Coupons:        coupons,
		Credits:        row.Credits.String(),
		CreditsPerUSD:  row.CreditsPerUsd.String(),
		MaxConsecutive: row.MaxConsecutive,
		InviterLayers:  row.InviterLayers,
		GoodID:         row.GoodID.String(),
		CreatedAt:      row.CreatedAt,
		UpdatedAt:      row.UpdatedAt,
	}

	return ret
}

func Ent2GrpcMany(rows []*ent.Event) []*npool.Event {
	infos := []*npool.Event{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
