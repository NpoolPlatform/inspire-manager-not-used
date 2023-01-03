package invitationcode

import (
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/invitationcode"
)

func Ent2Grpc(row *ent.InvitationCode) *npool.InvitationCode {
	if row == nil {
		return nil
	}

	return &npool.InvitationCode{
		ID:             row.ID.String(),
		AppID:          row.AppID.String(),
		UserID:         row.UserID.String(),
		InvitationCode: row.InvitationCode,
		Confirmed:      row.Confirmed,
		Disabled:       row.Disabled,
		CreatedAt:      row.CreatedAt,
		UpdatedAt:      row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.InvitationCode) []*npool.InvitationCode {
	infos := []*npool.InvitationCode{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
