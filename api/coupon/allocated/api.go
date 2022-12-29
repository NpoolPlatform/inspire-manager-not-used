package allocated

import (
	"github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	allocated.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	allocated.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
