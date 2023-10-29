package errutil

import (
	"context"
	"net/http"

	ctxutil "github.com/tyeryan/l-protocol/context"
	"google.golang.org/grpc/codes"
	"gopkg.in/go-playground/validator.v8"
)

var (
	grpcCodeToHTTPStatusCode = map[codes.Code]int{
		//400
		codes.InvalidArgument:    http.StatusBadRequest, //used
		codes.NotFound:           http.StatusBadRequest,
		codes.AlreadyExists:      http.StatusBadRequest,
		codes.FailedPrecondition: http.StatusBadRequest,
		codes.OutOfRange:         http.StatusBadRequest,
		codes.Canceled:           http.StatusBadRequest, //used, rate limiter
		codes.ResourceExhausted:  http.StatusBadRequest,
		codes.PermissionDenied:   http.StatusBadRequest, //used, deviec checking

		//401
		codes.Unauthenticated: http.StatusUnauthorized,

		//408
		codes.DeadlineExceeded: http.StatusRequestTimeout,

		//500 not retry
		codes.DataLoss: http.StatusInternalServerError,
		codes.Aborted:  http.StatusInternalServerError, //used, main logic finished, but additional action failed

		//500 retry
		codes.Internal:      http.StatusInternalServerError,
		codes.Unavailable:   http.StatusInternalServerError,
		codes.Unknown:       http.StatusInternalServerError,
		codes.Unimplemented: http.StatusInternalServerError,
	}

	defaultHttpErrorMap = map[ErrorCode]HTTPError{
		DefaultError: {
			Code:       InternalServerError,
			Desc:       "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
		},
		InvalidParamError: {
			Code:       InvalidParamError,
			Desc:       "the request parameter does not pass the validation checking",
			StatusCode: http.StatusBadRequest,
		},
		MfaResourcesDenied: {
			Code:       MfaResourcesDenied,
			Desc:       "User not allowed to access the requested resources without MFA",
			StatusCode: http.StatusUnauthorized,
		},
	}
)

type HTTPError struct {
	StatusCode int       `json:"-"`
	RequestID  string    `json:"request_id"`
	Code       ErrorCode `json:"code"`
	Desc       string    `json:"desc"`
}

func (e *HTTPError) ToGinJson() (int, *HTTPError) {
	return e.StatusCode, e
}

func toHttpStatus(grpcCode codes.Code) int {

	if statusCode, ok := grpcCodeToHTTPStatusCode[grpcCode]; ok {
		return statusCode
	}
	return http.StatusInternalServerError
}

// ToHTTPError convert to HTTP Error
func (r *RPCError) ToHTTPError() *HTTPError {
	if len(r.Detail) > 0 {
		detail := r.Detail[(len(r.Detail) - 1)]
		return &HTTPError{
			StatusCode: toHttpStatus(detail.GrpcCode),
			RequestID:  detail.Stan,
			Code:       detail.Code,
			Desc:       detail.Desc,
		}
	} else {
		httpError := defaultHttpErrorMap[DefaultHTTPError]
		return &httpError
	}

}

// ParseParamValidationError parse validation error to http error
func ParseParamValidationError(ctx context.Context, err error) *HTTPError {

	result, _ := GetHTTPError(ctx, InvalidParamError)
	err, isValidationError := err.(validator.ValidationErrors)
	if isValidationError {
		result.Desc = err.Error()
	} else {
		log.Errore(ctx, "not validation error", err)
	}
	return result
}

// ParseRPCErrorToHTTP parse rpc error to http error
func ParseRPCErrorToHTTP(ctx context.Context, srcErr error) *HTTPError {
	rpcErr := ParseRPCErrorRaw(ctx, srcErr)
	httpErr := rpcErr.ToHTTPError()
	if httpErr.StatusCode == http.StatusInternalServerError {
		log.Errorw(ctx, "Internal Server Error", "rpcErr", rpcErr)
	}
	return httpErr
}

// GetHTTPError Get predefined error
func GetHTTPError(ctx context.Context, code ErrorCode) (*HTTPError, bool) {
	requestID, _ := ctxutil.Read(ctx, ctxutil.Stan)
	result, ok := defaultHttpErrorMap[code]
	if !ok {
		result = defaultHttpErrorMap[DefaultError]
	}
	result.RequestID = requestID
	return &result, ok
}

func addHttpErrorDefs(httpErrors []*HTTPError) {
	for _, httpError := range httpErrors {
		defaultHttpErrorMap[httpError.Code] = *httpError
	}
}

func AddHttpErrorDefs(httpErrors []*HTTPError) {
	addHttpErrorDefs(httpErrors)
}
