package controller

import (
	"net/http"

	cmerror "victorzhou123/vicblog/common/domain/error"
)

const (
	errorSystemError     = "system_error"
	errorBadRequestBody  = "bad_request_body"
	errorBadRequestParam = "bad_request_param"
)

type errorCode interface {
	ErrorCode() string
}

type errorNotFound interface {
	errorCode

	NotFound()
}

type errorNoPermission interface {
	errorCode

	NoPermission()
}

func httpError(err error) (int, string) {
	if err == nil {
		return http.StatusOK, ""
	}

	sc := http.StatusInternalServerError
	code := errorSystemError

	if v, ok := err.(errorCode); ok {
		code = v.ErrorCode()

		if _, ok := err.(errorNotFound); ok {
			sc = http.StatusNotFound

		} else if _, ok := err.(errorNoPermission); ok {
			sc = http.StatusForbidden

		} else {
			switch code {
			case cmerror.ErrorCodeTokenInvalid:
				sc = http.StatusUnauthorized

			case cmerror.ErrorCodeAccessCertificateInvalid:
				sc = http.StatusUnauthorized

			default:
				sc = http.StatusBadRequest
			}
		}
	}

	return sc, code
}
