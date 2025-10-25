package validator

import (
	"errors"
	"strings"

	idLocales "github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/id"
	"golang.org/x/exp/constraints"
)

func ValidateStruct(obj interface{}) error {
	var errx []string

	idn := idLocales.New()
	uni := ut.New(idn, idn)
	transId, _ := uni.GetTranslator("id")
	validate := validator.New()

	_ = id.RegisterDefaultTranslations(validate, transId)
	if errs, ok := validate.Struct(obj).(validator.ValidationErrors); ok {
		for _, err := range errs {
			errx = append(errx, err.Translate(transId))
		}
	}

	if len(errx) > 0 {
		return errors.New(strings.Join(errx, ";"))
	}

	return nil
}

func IsNotNilAndNotEmptyString(val *string) bool {
	return val != nil && strings.TrimSpace(*val) != ""
}

func IsNotEmptyString(val string) bool {
	return strings.TrimSpace(val) != ""
}

func IsNotNilAndNotZeroInt[T constraints.Integer](val *T) bool {
	return val != nil && *val != 0
}

func IsNotNilAndNotZeroInt64(val *int64) bool {
	return val != nil && *val != 0
}

func IsNotNilAndNotZeroInt32(val *int32) bool {
	return val != nil && *val != 0
}

func IsNotNilAndTrue(val *bool) bool {
	return val != nil && *val
}

func IsNotNilAndFalse(val *bool) bool {
	return val != nil && !*val
}

func IsNotEmptySlice[T any](val []T) bool {
	return len(val) != 0
}
