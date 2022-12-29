package discount

import (
	"github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/discount"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	discount.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	discount.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
