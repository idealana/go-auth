package validator

import (
    "reflect"
    "strings"
    "github.com/go-playground/validator/v10"
    "github.com/go-playground/locales/en"
    ut "github.com/go-playground/universal-translator"
    en_trans "github.com/go-playground/validator/v10/translations/en"
)

type ValidationError struct {
    Errors map[string]string
}

func (v *ValidationError) Error() string {
    return "Validation error."
}

type structValidator struct {
    validate *validator.Validate
    trans ut.Translator
}

func (v *structValidator) Validate(out any) error {
    err := v.validate.Struct(out)

    if err == nil {
        return nil
    }

    validationErrors := err.(validator.ValidationErrors)
    errors := make(map[string]string)

    for _, e := range validationErrors {
        errors[e.Field()] = e.Translate(v.trans)
    }

    return &ValidationError{Errors: errors}
}

func NewValidator() *structValidator {
    validate := validator.New()

    validate.RegisterTagNameFunc(func(field reflect.StructField) string {
        name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]

        if name == "-" {
            return ""
        }

        return name
    })

    en := en.New()
    uni := ut.New(en, en)

    trans, _ := uni.GetTranslator("en")

    en_trans.RegisterDefaultTranslations(validate, trans)

    return &structValidator{
        validate: validate,
        trans: trans,
    }
}
