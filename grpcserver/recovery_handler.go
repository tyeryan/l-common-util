package grpcserver

import (
	"context"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/pkg/errors"
	"github.com/tyeryan/l-protocol/errutil"
	"runtime/debug"
)

func GRPCErrorWrapper() grpc_recovery.RecoveryHandlerFuncContext {
	return func(ctx context.Context, p interface{}) (err error) {
		log.Errore(ctx, "grpc recover from panic", errors.WithStack(errors.Errorf("%s", p))) // until we can propagate the stack in grpc error, we need to explicitly log it
		log.Errorw(ctx, "panic stack trace", string(debug.Stack()))                          // print stacktrace as above doesn't

		return errutil.NewRPCError(ctx, errors.WithStack(errors.Errorf("%s", p)))
	}
}
