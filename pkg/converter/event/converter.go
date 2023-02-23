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

	ret := &npool.Event{
		ID:             row.ID.String(),
		AppID:          row.AppID.String(),
		EventType:      basetypes.UsedFor(basetypes.UsedFor_value[row.EventType]),
		Coupons:        row.Coupons,
		Credits:        row.Credits.String(),
		CreditsPerUSD:  row.CreditsPerUsd.String(),
		MaxConsecutive: row.MaxConsecutive,
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