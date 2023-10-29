package errutil

import (
	"google.golang.org/grpc/codes"
)

const (
	//InternalServerError default error response
	InternalServerError      ErrorCode = "INTERNAL_SERVER_ERROR"
	InvalidParamError        ErrorCode = "INVALID_PARAM_ERROR"
	TooManyRequest           ErrorCode = "TOO_MANY_REQUEST"
	InternalServerErrorAbort ErrorCode = "INTERNAL_SERVER_ERROR"
	AccessResourcesDenied    ErrorCode = "ACCESS_RESOURCESES_DENIED"
	MfaResourcesDenied       ErrorCode = "MFA_RESOURCES_DENIED"
	ResourceNotFound         ErrorCode = "RESOURCE_NOT_FOUND"
	ServiceUnavailable       ErrorCode = "SERVICE_UNAVAILABLE"
	RecordAlreadyExists      ErrorCode = "RECORD_ALREADY_EXISTS"
)

var (
	//rpc error map init
	templates = []*Template{
		//add default error
		{
			Code:     InternalServerError,
			Desc:     "Internal Service Error",
			GrpcCode: codes.Internal,
		},
		{
			Code:     InvalidParamError,
			Desc:     "The request parameter does not pass the validation checking",
			GrpcCode: codes.Internal,
		},
		{
			Code:     TooManyRequest,
			Desc:     "Too many request, please wait and try later",
			GrpcCode: codes.ResourceExhausted,
		},
		{
			Code:     InternalServerErrorAbort,
			Desc:     "Internal Service Error",
			GrpcCode: codes.Aborted,
		},
		{
			Code:     AccessResourcesDenied,
			Desc:     "User does not allowed to access the requested resources",
			GrpcCode: codes.PermissionDenied,
		},
		{
			Code:     MfaResourcesDenied,
			Desc:     "User not allowed to access the requested resources without MFA",
			GrpcCode: codes.PermissionDenied,
		},
		{
			Code:     ResourceNotFound,
			Desc:     "Resource Not Found",
			GrpcCode: codes.NotFound,
		},
		{
			Code:     ServiceUnavailable,
			Desc:     "The server cannot handle the request (because it is overloaded or down for maintenance)",
			GrpcCode: codes.Unavailable,
		},
		{
			Code:     RecordAlreadyExists,
			Desc:     "Record already exits, Duplication check failed",
			GrpcCode: codes.Internal,
		},
	}
)

func init() {
	addErrorTemplate(templates)
}
