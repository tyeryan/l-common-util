package errutil

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	ctxutil "github.com/tyeryan/l-protocol/context"
	pb "github.com/tyeryan/l-protocol/go/errors"
	llog "github.com/tyeryan/l-protocol/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"runtime"
)

var log = llog.GetLogger("error-util")

func NewRPCError(ctx context.Context, err error, opts ...Option) error {
	//set option
	option := &options{}
	for _, opt := range opts {
		opt(option)
	}

	//use option to update template
	errT := getErrorOrDefault(ctx, option.Code) //option code
	if option.Desc != "" {                      //option desc
		errT.Desc = option.Desc
	}

	//get stan
	stan, _ := ctxutil.Read(ctx, ctxutil.Stan)

	//create a new status error
	curStatus := status.New(errT.GrpcCode, err.Error())

	//if has source, generate a status error from source
	if option.Source != nil { //option source
		preStatus, ok := status.FromError(option.Source)
		if ok {
			for _, detail := range preStatus.Details() {
				if msg, ok := detail.(proto.Message); ok {
					curStatus, _ = curStatus.WithDetails(msg)
				}
			}
		}
	}

	//create current error detail
	curDetail := &ErrorDetail{
		Code:       errT.Code,
		Desc:       errT.Desc,
		Error:      err.Error(),
		GrpcCode:   errT.GrpcCode,
		Stan:       stan,
		CallerInfo: callerInfo(ctx),
	}

	curStatus, _ = curStatus.WithDetails(curDetail.ToPB())
	return curStatus.Err()
}

func ParseRPCErrorRaw(ctx context.Context, source error) *RPCError {
	//get stan
	stan, _ := ctxutil.Read(ctx, ctxutil.Stan)

	statusErr, ok := status.FromError(source)
	if !ok {
		log.Alertw(ctx, fmt.Sprintf("error is not RPCError, please check your implementation\nSource Error:%v\n", source))
		defaultTemplate := getErrorOrDefault(ctx, DefaultError)
		return &RPCError{
			Code:  codes.Internal,
			Error: source,
			Detail: []*ErrorDetail{
				{
					Code:       defaultTemplate.Code,
					GrpcCode:   defaultTemplate.GrpcCode,
					Desc:       defaultTemplate.Desc,
					Error:      source.Error(),
					CallerInfo: callerInfo(ctx),
					Stan:       stan,
				},
			},
		}
	}
	rpcErr := newRPCError(ctx, statusErr)
	return rpcErr
}

func newRPCError(ctx context.Context, s *status.Status) *RPCError {
	e := &RPCError{
		Code:  s.Code(),
		Error: errors.New(s.Message()),
	}

	for _, detail := range s.Details() {
		if pbDetail, ok := detail.(*pb.ErrorDetail); ok {
			detail := ErrorDetailFromPB(pbDetail)
			e.Detail = append(e.Detail, detail)
		}
	}
	return e
}

func callerInfo(ctx context.Context) *CallerInfo {

	serviceName := os.Getenv("SERVICECONFIG_SERVICENAME")
	if serviceName == "" {
		log.Alertw(ctx, "serviceName is empty, please check configuration")
	}

	fpcs := make([]uintptr, 3)
	n := runtime.Callers(3, fpcs)
	if n == 0 {
		return nil
	}

	fun := runtime.FuncForPC(fpcs[0] - 1)
	if fun == nil {
		return nil
	}
	file, line := fun.FileLine(fpcs[0])
	return &CallerInfo{
		ServiceName: serviceName,
		FuncName:    fun.Name(),
		FileName:    file,
		LineNum:     int32(line),
	}
}
