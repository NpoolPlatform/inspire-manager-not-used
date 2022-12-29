package goodorderpercent

import (
	"github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission/goodorderpercent"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	goodorderpercent.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	goodorderpercent.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
