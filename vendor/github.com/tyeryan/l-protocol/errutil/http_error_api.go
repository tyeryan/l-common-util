package errutil

import (
	"net/http"
)

const (
	DefaultHTTPError   ErrorCode = "INTERNAL_SERVER_ERROR"
	InvalidClientAuth  ErrorCode = "INVALID_CLIENT_AUTH"
	InvalidAccessToken ErrorCode = "INVALID_ACCESS_TOKEN"
	NoPermissionAPI    ErrorCode = "NO_PERMISSION_API"
	MissingRequestID   ErrorCode = "MISSING_REQUEST_ID"
	InvalidCaptcha     ErrorCode = "INVALID_CAPTCHA_ERROR"
)

func init() {
	httpErrors := []*HTTPError{
		{
			Code:       InvalidAccessToken,
			Desc:       "The access token is invalid or expired",
			StatusCode: http.StatusUnauthorized,
		},
		{
			Code:       InvalidClientAuth,
			Desc:       "The client auth credential is invalid",
			StatusCode: http.StatusUnauthorized,
		},
		{
			Code:       NoPermissionAPI,
			Desc:       "the access token does not have the scope to access the requested API",
			StatusCode: http.StatusUnauthorized,
		},
		{
			Code:       MissingRequestID,
			Desc:       "The request id is missing",
			StatusCode: http.StatusBadRequest,
		},
		{
			Code:       InternalServerError,
			Desc:       "Internal Service Error",
			StatusCode: http.StatusInternalServerError,
		},
		{
			Code:       InvalidCaptcha,
			Desc:       "Invalid Captcha Error",
			StatusCode: http.StatusUnauthorized,
		},
	}
	addHttpErrorDefs(httpErrors)
}
