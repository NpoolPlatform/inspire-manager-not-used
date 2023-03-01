package event

import (
	"github.com/NpoolPlatform/message/npool/inspire/mgr/v1/event"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	event.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	event.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
