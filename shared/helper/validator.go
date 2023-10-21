package helper

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Message []string `json:"message"`
}

func BuildResponse(data interface{}) Response {
	return Response{
		Message: []string{},
	}
}

func ValidationErrResponse(err error) Response {
	var ve validator.ValidationErrors
	res := Response{
		Message: make([]string, 0),
	}

	if !errors.As(err, &ve) {
		return res
	}

	for _, r := range ve {
		res.Message = append(res.Message, msgForTag(r))
	}

	return res
}

func msgForTag(r validator.FieldError) string {
	switch r.Tag() {
	case "required":
		return fmt.Sprintf("%s field is a required field", strings.ToLower(r.Field()))
	case "email":
		return fmt.Sprintf("%s field is not a valid email", r.Field())
	case "min":
		return fmt.Sprintf("%s field minimun is %s", r.Field(), r.Param())
	case "max":
		return fmt.Sprintf("%s field maximum is %s", r.Field(), r.Param())
	case "oneof":
		return fmt.Sprintf("%s field should be one of %s", r.Field(), strings.Join(strings.Split(r.Param(), " "), ","))
	}
	return ""
}
