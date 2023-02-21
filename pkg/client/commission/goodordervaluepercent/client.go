//nolint:dupl
package goodordervaluepercent

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission/goodordervaluepercent"

	constant "github.com/NpoolPlatform/inspire-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get goodordervaluepercent connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateOrderValuePercent(ctx context.Context, in *npool.OrderValuePercentReq) (*npool.OrderValuePercent, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateOrderValuePercent(ctx, &npool.CreateOrderValuePercentRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create goodordervaluepercent: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create goodordervaluepercent: %v", err)
	}
	return info.(*npool.OrderValuePercent), nil
}

func CreateOrderValuePercents(ctx context.Context, in []*npool.OrderValuePercentReq) ([]*npool.OrderValuePercent, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateOrderValuePercents(ctx, &npool.CreateOrderValuePercentsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create goodordervaluepercents: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create goodordervaluepercents: %v", err)
	}
	return infos.([]*npool.OrderValuePercent), nil
}

func UpdateOrderValuePercent(ctx context.Context, in *npool.OrderValuePercentReq) (*npool.OrderValuePercent, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateOrderValuePercent(ctx, &npool.UpdateOrderValuePercentRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update goodordervaluepercent: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update goodordervaluepercent: %v", err)
	}
	return info.(*npool.OrderValuePercent), nil
}

func GetOrderValuePercent(ctx context.Context, id string) (*npool.OrderValuePercent, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetOrderValuePercent(ctx, &npool.GetOrderValuePercentRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get goodordervaluepercent: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get goodordervaluepercent: %v", err)
	}
	return info.(*npool.OrderValuePercent), nil
}

func GetOrderValuePercentOnly(ctx context.Context, conds *npool.Conds) (*npool.OrderValuePercent, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetOrderValuePercentOnly(ctx, &npool.GetOrderValuePercentOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get goodordervaluepercent: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get goodordervaluepercent: %v", err)
	}
	return info.(*npool.OrderValuePercent), nil
}

func GetOrderValuePercents(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.OrderValuePercent, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetOrderValuePercents(ctx, &npool.GetOrderValuePercentsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get goodordervaluepercents: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get goodordervaluepercents: %v", err)
	}
	return infos.([]*npool.OrderValuePercent), total, nil
}

func ExistOrderValuePercent(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistOrderValuePercent(ctx, &npool.ExistOrderValuePercentRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get goodordervaluepercent: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get goodordervaluepercent: %v", err)
	}
	return infos.(bool), nil
}

func ExistOrderValuePercentConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistOrderValuePercentConds(ctx, &npool.ExistOrderValuePercentCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get goodordervaluepercent: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get goodordervaluepercent: %v", err)
	}
	return infos.(bool), nil
}

func CountOrderValuePercents(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountOrderValuePercents(ctx, &npool.CountOrderValuePercentsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count goodordervaluepercent: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count goodordervaluepercent: %v", err)
	}
	return infos.(uint32), nil
}
