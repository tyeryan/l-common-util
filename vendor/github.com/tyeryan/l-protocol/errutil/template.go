package errutil

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
)

const (
	DefaultError = InternalServerError
)

var (
	errorMap = map[ErrorCode]*Template{}
)

// ErrorCode unique error code
type ErrorCode string

// Template error template
type Template struct {
	Code     ErrorCode
	GrpcCode codes.Code
	Desc     string
}

func getErrorOrDefault(ctx context.Context, code ErrorCode) *Template {

	if code == "" {
		code = DefaultError
	}

	dest := &Template{}
	var err *Template
	e, ok := errorMap[code]
	if ok {
		err = e
	} else {
		log.Alertw(ctx, fmt.Sprintf("error code not found:%s, please check the implementation.", code))
		err = errorMap[DefaultError]
	}

	//clone template
	copier.Copy(dest, err)
	return dest
}

// add error template array to error map
func addErrorTemplate(templates []*Template) {
	// When Template.Code is already mapped, it will be overwritten.
	// Meaning that a ErrorCode will map to the latest assigned Template's GrpcCode.
	//
	// Example:
	// 		InternalServerError      ErrorCode = "INTERNAL_SERVER_ERROR" => codes.Internal
	// 		InternalServerErrorAbort ErrorCode = "INTERNAL_SERVER_ERROR" => codes.Aborted
	// 		"INTERNAL_SERVER_ERROR" will map to codes.Aborted.
	for key, resp := range templates {
		errorMap[resp.Code] = templates[key]
	}
}

func AddErrorTemplate(templates []*Template) {
	addErrorTemplate(templates)
}

func NewTemplateFromErrorDetail(errDetail ErrorDetail) *Template {
	return &Template{
		Code:     errDetail.Code,
		GrpcCode: errDetail.GrpcCode,
		Desc:     errDetail.Desc,
	}
}
