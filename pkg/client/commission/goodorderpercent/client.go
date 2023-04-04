//nolint:dupl
package goodorderpercent

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission/goodorderpercent"

	"github.com/NpoolPlatform/inspire-manager/pkg/servicename"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(servicename.ServiceDomain, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get goodorderpercent connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateOrderPercent(ctx context.Context, in *npool.OrderPercentReq) (*npool.OrderPercent, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateOrderPercent(ctx, &npool.CreateOrderPercentRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create goodorderpercent: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create goodorderpercent: %v", err)
	}
	return info.(*npool.OrderPercent), nil
}

func CreateOrderPercents(ctx context.Context, in []*npool.OrderPercentReq) ([]*npool.OrderPercent, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateOrderPercents(ctx, &npool.CreateOrderPercentsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create goodorderpercents: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create goodorderpercents: %v", err)
	}
	return infos.([]*npool.OrderPercent), nil
}

func UpdateOrderPercent(ctx context.Context, in *npool.OrderPercentReq) (*npool.OrderPercent, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateOrderPercent(ctx, &npool.UpdateOrderPercentRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update goodorderpercent: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update goodorderpercent: %v", err)
	}
	return info.(*npool.OrderPercent), nil
}

func GetOrderPercent(ctx context.Context, id string) (*npool.OrderPercent, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetOrderPercent(ctx, &npool.GetOrderPercentRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get goodorderpercent: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get goodorderpercent: %v", err)
	}
	return info.(*npool.OrderPercent), nil
}

func GetOrderPercentOnly(ctx context.Context, conds *npool.Conds) (*npool.OrderPercent, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetOrderPercentOnly(ctx, &npool.GetOrderPercentOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get goodorderpercent: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get goodorderpercent: %v", err)
	}
	return info.(*npool.OrderPercent), nil
}

func GetOrderPercents(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.OrderPercent, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetOrderPercents(ctx, &npool.GetOrderPercentsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get goodorderpercents: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get goodorderpercents: %v", err)
	}
	return infos.([]*npool.OrderPercent), total, nil
}

func ExistOrderPercent(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistOrderPercent(ctx, &npool.ExistOrderPercentRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get goodorderpercent: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get goodorderpercent: %v", err)
	}
	return infos.(bool), nil
}

func ExistOrderPercentConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistOrderPercentConds(ctx, &npool.ExistOrderPercentCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get goodorderpercent: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get goodorderpercent: %v", err)
	}
	return infos.(bool), nil
}

func CountOrderPercents(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountOrderPercents(ctx, &npool.CountOrderPercentsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count goodorderpercent: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count goodorderpercent: %v", err)
	}
	return infos.(uint32), nil
}
