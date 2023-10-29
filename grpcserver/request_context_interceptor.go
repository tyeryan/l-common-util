package grpcserver

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func RequestContextUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return handler(chainContext(ctx), req)
}

func chainContext(ctx context.Context) context.Context {
	inMD, _ := metadata.FromIncomingContext(ctx)
	newCtx := metadata.NewOutgoingContext(ctx, inMD)
	return newCtx
}
