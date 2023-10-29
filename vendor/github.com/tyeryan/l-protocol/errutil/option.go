package errutil

// Option option function
type Option func(*options)

type options struct {
	Code   ErrorCode
	Desc   string
	Source error
}

// WithSource with source error
func WithSource(source error) Option {
	return func(opt *options) {
		if source != nil {
			opt.Source = source
		}
	}
}

// WithCode replace the default error code by parameter error code
func WithCode(code ErrorCode) Option {
	return func(opt *options) {
		if code != "" {
			opt.Code = code
		}
	}
}

// WithDesc replace the default description by paramter description
func WithDesc(desc string) Option {
	return func(opt *options) {
		if desc != "" {
			opt.Desc = desc
		}
	}
}
