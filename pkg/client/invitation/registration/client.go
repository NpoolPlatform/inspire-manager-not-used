//nolint:dupl
package registration

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"

	"github.com/NpoolPlatform/inspire-manager/pkg/servicename"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(servicename.ServiceDomain, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get registration connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateRegistration(ctx context.Context, in *npool.RegistrationReq) (*npool.Registration, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateRegistration(ctx, &npool.CreateRegistrationRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create registration: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create registration: %v", err)
	}
	return info.(*npool.Registration), nil
}

func CreateRegistrations(ctx context.Context, in []*npool.RegistrationReq) ([]*npool.Registration, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateRegistrations(ctx, &npool.CreateRegistrationsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create registrations: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create registrations: %v", err)
	}
	return infos.([]*npool.Registration), nil
}

func UpdateRegistration(ctx context.Context, in *npool.RegistrationReq) (*npool.Registration, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateRegistration(ctx, &npool.UpdateRegistrationRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update registration: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update registration: %v", err)
	}
	return info.(*npool.Registration), nil
}

func GetRegistration(ctx context.Context, id string) (*npool.Registration, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetRegistration(ctx, &npool.GetRegistrationRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get registration: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get registration: %v", err)
	}
	return info.(*npool.Registration), nil
}

func GetRegistrationOnly(ctx context.Context, conds *npool.Conds) (*npool.Registration, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetRegistrationOnly(ctx, &npool.GetRegistrationOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get registration: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get registration: %v", err)
	}
	return info.(*npool.Registration), nil
}

func GetRegistrations(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Registration, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetRegistrations(ctx, &npool.GetRegistrationsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get registrations: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get registrations: %v", err)
	}
	return infos.([]*npool.Registration), total, nil
}

func ExistRegistration(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistRegistration(ctx, &npool.ExistRegistrationRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get registration: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get registration: %v", err)
	}
	return infos.(bool), nil
}

func ExistRegistrationConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistRegistrationConds(ctx, &npool.ExistRegistrationCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get registration: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get registration: %v", err)
	}
	return infos.(bool), nil
}

func CountRegistrations(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountRegistrations(ctx, &npool.CountRegistrationsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count registration: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count registration: %v", err)
	}
	return infos.(uint32), nil
}
