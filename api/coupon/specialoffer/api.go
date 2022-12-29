package specialoffer

import (
	"github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/specialoffer"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	specialoffer.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	specialoffer.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
