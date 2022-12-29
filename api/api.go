package api

import (
	"context"

	"github.com/NpoolPlatform/inspire-manager/api/archivement/detail"
	"github.com/NpoolPlatform/inspire-manager/api/archivement/general"
	"github.com/NpoolPlatform/inspire-manager/api/commission/goodorderpercent"
	"github.com/NpoolPlatform/inspire-manager/api/coupon/allocated"
	"github.com/NpoolPlatform/inspire-manager/api/coupon/discount"
	"github.com/NpoolPlatform/inspire-manager/api/coupon/fixamount"
	"github.com/NpoolPlatform/inspire-manager/api/coupon/specialoffer"
	"github.com/NpoolPlatform/inspire-manager/api/invitation/invitationcode"
	"github.com/NpoolPlatform/inspire-manager/api/invitation/registration"

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
	general.Register(server)
	detail.Register(server)
	invitationcode.Register(server)
	registration.Register(server)
	goodorderpercent.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := v1.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
