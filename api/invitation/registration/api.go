package registration

import (
	"github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	registration.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	registration.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
