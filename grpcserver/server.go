package grpcserver

import (
	"context"
	"github.com/google/wire"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.elastic.co/apm/module/apmgrpc"

	"github.com/tyeryan/l-common-util/apm"
	logutil "github.com/tyeryan/l-protocol/log"
	"google.golang.org/grpc"
	"net"
)

var (
	WireSet = wire.NewSet(
		ProvideServer,
	)
	log              = logutil.GetLogger("grpc-server")
	messageSizeLimit = 30 * 1024 * 1024
)

type Server struct {
	GPRCServer *grpc.Server
	listener   net.Listener
}

func ProvideServer(ctx context.Context, apmConfig *apm.ApmConfig) (*Server, error) {
	port := ":8888"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}

	grpcServer := grpc.NewServer(
		grpc.MaxSendMsgSize(messageSizeLimit),
		grpc.MaxRecvMsgSize(messageSizeLimit),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandlerContext(GRPCErrorWrapper())),
		)),
	)
	return &Server{
		GPRCServer: grpcServer,
		listener:   lis,
	}, nil
}

func ProvideUnaryServerInterceptor(apmConfig *apm.ApmConfig) grpc.UnaryServerInterceptor {
	if apmConfig.Enable {
		return grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandlerContext(GRPCErrorWrapper())),
			RequestContextUnaryInterceptor,
			apmgrpc.NewUnaryServerInterceptor(),
		)
	} else {
		return grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandlerContext(GRPCErrorWrapper())),
			RequestContextUnaryInterceptor,
		)
	}

}

// Serve accepts incoming connections on the listener lis, creating a new
// ServerTransport and service goroutine for each. The service goroutines
// read gRPC requests and then call the registered handlers to reply to them.
// Serve returns when lis.Accept fails with fatal errors.  lis will be closed when
// this method returns.
// Serve will return a non-nil error unless Stop or GracefulStop is called.
func (s *Server) Serve() error {
	return s.GPRCServer.Serve(s.listener)
}
