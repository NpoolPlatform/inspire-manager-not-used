//nolint:dupl
package invitationcode

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/invitationcode"

	"github.com/NpoolPlatform/inspire-manager/pkg/servicename"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(servicename.ServiceDomain, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get invitationcode connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateInvitationCode(ctx context.Context, in *npool.InvitationCodeReq) (*npool.InvitationCode, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateInvitationCode(ctx, &npool.CreateInvitationCodeRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create invitationcode: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create invitationcode: %v", err)
	}
	return info.(*npool.InvitationCode), nil
}

func CreateInvitationCodes(ctx context.Context, in []*npool.InvitationCodeReq) ([]*npool.InvitationCode, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateInvitationCodes(ctx, &npool.CreateInvitationCodesRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create invitationcodes: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create invitationcodes: %v", err)
	}
	return infos.([]*npool.InvitationCode), nil
}

func UpdateInvitationCode(ctx context.Context, in *npool.InvitationCodeReq) (*npool.InvitationCode, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateInvitationCode(ctx, &npool.UpdateInvitationCodeRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update invitationcode: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update invitationcode: %v", err)
	}
	return info.(*npool.InvitationCode), nil
}

func GetInvitationCode(ctx context.Context, id string) (*npool.InvitationCode, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetInvitationCode(ctx, &npool.GetInvitationCodeRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get invitationcode: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get invitationcode: %v", err)
	}
	return info.(*npool.InvitationCode), nil
}

func GetInvitationCodeOnly(ctx context.Context, conds *npool.Conds) (*npool.InvitationCode, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetInvitationCodeOnly(ctx, &npool.GetInvitationCodeOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get invitationcode: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get invitationcode: %v", err)
	}
	return info.(*npool.InvitationCode), nil
}

func GetInvitationCodes(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.InvitationCode, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetInvitationCodes(ctx, &npool.GetInvitationCodesRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get invitationcodes: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get invitationcodes: %v", err)
	}
	return infos.([]*npool.InvitationCode), total, nil
}

func ExistInvitationCode(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistInvitationCode(ctx, &npool.ExistInvitationCodeRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get invitationcode: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get invitationcode: %v", err)
	}
	return infos.(bool), nil
}

func ExistInvitationCodeConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistInvitationCodeConds(ctx, &npool.ExistInvitationCodeCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get invitationcode: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get invitationcode: %v", err)
	}
	return infos.(bool), nil
}

func CountInvitationCodes(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountInvitationCodes(ctx, &npool.CountInvitationCodesRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count invitationcode: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count invitationcode: %v", err)
	}
	return infos.(uint32), nil
}
