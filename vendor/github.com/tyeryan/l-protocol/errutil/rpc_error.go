package errutil

import (
	pb "github.com/tyeryan/l-protocol/go/errors"
	"google.golang.org/grpc/codes"
)

// RPCError rpc error
type RPCError struct {
	Code   codes.Code
	Error  error
	Detail []*ErrorDetail
}

// ErrorDetail error detail dto
type ErrorDetail struct {
	Stan       string
	Code       ErrorCode
	GrpcCode   codes.Code
	Desc       string
	Error      string
	CallerInfo *CallerInfo
}

// CallerInfo caller info for debuging
type CallerInfo struct {
	ServiceName string
	FuncName    string
	FileName    string
	LineNum     int32
}

// CallerInfoFromPB create caller info from protocol buffer
func CallerInfoFromPB(msg *pb.Caller) *CallerInfo {
	return &CallerInfo{
		ServiceName: msg.ServiceName,
		FuncName:    msg.FuncName,
		FileName:    msg.FileName,
		LineNum:     msg.LineNum,
	}
}

// ToPB convert to protocol buffer
func (info *CallerInfo) ToPB() *pb.Caller {
	return &pb.Caller{
		ServiceName: info.ServiceName,
		FuncName:    info.FuncName,
		FileName:    info.FileName,
		LineNum:     info.LineNum,
	}
}

func ErrorDetailFromPB(msg *pb.ErrorDetail) *ErrorDetail {
	return &ErrorDetail{
		Code:       ErrorCode(msg.Code),
		Desc:       msg.Desc,
		Error:      msg.Error,
		GrpcCode:   codes.Code(msg.GrpcCode),
		Stan:       msg.Stan,
		CallerInfo: CallerInfoFromPB(msg.CallerInfo),
	}
}

// ToPB convert to protocol buffer
func (e *ErrorDetail) ToPB() *pb.ErrorDetail {
	return &pb.ErrorDetail{
		Code:       string(e.Code),
		Desc:       e.Desc,
		Error:      e.Error,
		GrpcCode:   uint32(e.GrpcCode),
		Stan:       e.Stan,
		CallerInfo: e.CallerInfo.ToPB(),
	}
}
