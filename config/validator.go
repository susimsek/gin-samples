package config

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	"reflect"
	"strings"
)

func NewValidator() (*validator.Validate, ut.Translator) {
	validate := validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	translator := en.New()
	uni := ut.New(translator, translator)
	trans, _ := uni.GetTranslator("en")

	entranslations.RegisterDefaultTranslations(validate, trans)

	registerCustomTranslations(validate, trans)

	return validate, trans
}

func registerCustomTranslations(validate *validator.Validate, trans ut.Translator) {
	// Required
	_ = validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "Field cannot be null", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	// Min Length
	_ = validate.RegisterTranslation("min", trans, func(ut ut.Translator) error {
		return ut.Add("min", "Field must be at least {1} characters", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", fe.Field(), fe.Param())
		return t
	})

	// Max Length
	_ = validate.RegisterTranslation("max", trans, func(ut ut.Translator) error {
		return ut.Add("max", "Field  must not exceed {1} characters", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("max", fe.Field(), fe.Param())
		return t
	})
}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T(fe.Tag(), fe.Field(), fe.Param())
	return t
}
