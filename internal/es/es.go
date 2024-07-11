package es

import (
	"github.com/go-kratos/kratos/v2/errors"
)

func NewHTTPError(code int, field string, detail string) *HTTPError {
	return &HTTPError{
		Errors: map[string][]string{
			field: []string{detail},
		},
		Code: code,
	}
}

type HTTPError struct {
	Errors map[string][]string `json:"errors"`
	Code   int                 `json:"-"`
}

func (e *HTTPError) Error() string {
	return e.Errors["body"][0]
}
func FromError(err error) *HTTPError {
	if err == nil {
		return nil
	}
	if se := new(HTTPError); errors.As(err, &se) {
		return se
	}

	if se := new(errors.Error); errors.As(err, &se) {
		return NewHTTPError(int(se.Code), se.Reason, se.Message)
	}
	return NewHTTPError(500, "internal", err.Error())
}
