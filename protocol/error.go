package protocol

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrReason int

const (
	ErrSample ErrReason = iota + 10001
)

var reasonMap = map[ErrReason]string{
	ErrSample: "This is a sample error",
}

type CustomError struct {
	Reason string
}

func (e *CustomError) Error() string {
	return e.Reason
}

func NewCustomError(reason string) *CustomError {
	return &CustomError{Reason: reason}
}

type ServerError struct {
	Code   ErrReason `json:"code"`
	Reason string    `json:"reason"`
}

func ErrorJSON(ctx *gin.Context, status int, errCode ErrReason, err error) {
	var reason string
	var verr validator.ValidationErrors
	var customError *CustomError
	if errors.As(err, &verr) && len(verr) > 0 {
		reason = fmt.Sprintf("Validate failed on field %s due to the rule '%s'", verr[0].Field(), verr[0].Tag())
	} else if errors.As(err, &customError) {
		reason = err.Error()
	} else {
		reason = reasonMap[errCode]
	}

	ctx.JSON(status, ServerError{
		Code:   errCode,
		Reason: reason,
	})
}

func ErrorString(errCode ErrReason) string {
	err := ServerError{
		Code:   errCode,
		Reason: reasonMap[errCode],
	}
	data, _ := json.Marshal(err)
	return string(data)
}
