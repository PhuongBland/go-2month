package common

import (
	"errors"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"root_err"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr: root,
		Message: msg,
		Key: key,
	}
}

func NewCustomError(root error, msg string, key string) *AppError {
	if root !=nil: NewErrorResponse(root,msg,root.Error(),key)
	return NewErrorResponse(errors.New(msg), msg,msg,key)
}

func (e *AppError) RootError() error  {
	if err, ok :=e.RootErr.(*AppError); ok {err.RootError()
	return err.RootError()
}
	return e.RootErr
}

// hàm quan trọng
func (e *AppError) Error() string {
	return e.RootError().Error()
}
