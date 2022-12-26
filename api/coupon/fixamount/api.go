package fixamount

import (
	"github.com/NpoolPlatform/message/npool/inspire/mgr/v1/inspire/coupon/fixamount"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	fixamount.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	fixamount.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
