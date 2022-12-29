package invitationcode

import (
	"github.com/NpoolPlatform/message/npool/inspire/mgr/v1/inspire/invitation/invitationcode"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	invitationcode.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	invitationcode.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
