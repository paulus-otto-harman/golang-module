package gola

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

type Error struct {
	Field string      `json:"field"`
	Tag   string      `json:"tag"`
	Value interface{} `json:"value"`
	Param string      `json:"param"`
}

func Validate(input interface{}) interface{} {
	validate := validator.New()
	validate.RegisterTagNameFunc(JsonFieldName)
	err := validate.Struct(input)

	if err != nil {
		var errorBags []Error
		validationErrors, hasValidationError := err.(validator.ValidationErrors)
		if hasValidationError {
			for _, inputError := range validationErrors {
				errorBags = append(errorBags, initError(inputError).ShouldHaveJsonParam(input))
			}
			return errorBags
		}
		return errors.New(fmt.Sprintf("validation error %v", err))
	}
	return nil
}

func (err *Error) ShouldHaveJsonParam(input interface{}) Error {
	if err.Tag == "eqfield" {
		if param, validationHasParam := reflect.TypeOf(input).FieldByName(err.Param); validationHasParam {
			err.Param = param.Tag.Get("json")
		}
	}
	return *err
}

func initError(inputError validator.FieldError) *Error {
	return &Error{
		Field: inputError.Field(),
		Tag:   inputError.Tag(),
		Value: inputError.Value(),
		Param: inputError.Param(),
	}
}

func JsonFieldName(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}
