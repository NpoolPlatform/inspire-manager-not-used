package api

import (
	"context"

	"github.com/NpoolPlatform/inspire-manager/api/coupon/allocated"
	"github.com/NpoolPlatform/inspire-manager/api/coupon/discount"
	"github.com/NpoolPlatform/inspire-manager/api/coupon/fixamount"
	"github.com/NpoolPlatform/inspire-manager/api/coupon/specialoffer"

	v1 "github.com/NpoolPlatform/message/npool/inspire/mgr/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	v1.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	v1.RegisterManagerServer(server, &Server{})
	allocated.Register(server)
	discount.Register(server)
	fixamount.Register(server)
	specialoffer.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := v1.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
