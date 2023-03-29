package goodordervaluepercent

import (
	"github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission/goodordervaluepercent"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	goodordervaluepercent.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	goodordervaluepercent.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
