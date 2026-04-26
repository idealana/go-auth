package validator

import (
    "fmt"
    "reflect"
    "strings"

    "github.com/go-playground/validator/v10"
    "github.com/go-playground/locales/en"
    "github.com/go-playground/locales/id"
    ut "github.com/go-playground/universal-translator"
    en_trans "github.com/go-playground/validator/v10/translations/en"
    id_trans "github.com/go-playground/validator/v10/translations/id"
)

const defaultLocale = "en"

type Validator interface {
    Validate(out any) error
    ValidateWithLocale(out any, locale string) error
}

type ValidationError struct {
    Errors map[string]string
}

func (v *ValidationError) Error() string {
    return "validation error"
}

type structValidator struct {
    validate *validator.Validate
    trans *ut.UniversalTranslator
}

func (v *structValidator) Validate(out any) error {
    return v.ValidateWithLocale(out, defaultLocale) // default locale
}

func (v *structValidator) ValidateWithLocale(out any, locale string) error {
    translator, found := v.trans.GetTranslator(locale)
    if !found {
        translator = v.trans.GetFallback() // fallback to default (en)
    }

    err := v.validate.Struct(out)
    if err == nil {
        return nil
    }

    validationErrors, ok := err.(validator.ValidationErrors)
    if !ok {
        return err
    }

    errors := make(map[string]string, len(validationErrors))
    for _, e := range validationErrors {
        errors[e.Field()] = e.Translate(translator)
    }

    return &ValidationError{Errors: errors}
}

func NewValidator() (Validator, error) {
    validate := validator.New()

    validate.RegisterTagNameFunc(func(field reflect.StructField) string {
        name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]

        if name == "-" {
            return ""
        }

        return name
    })

    enLocale := en.New()
    idLocale := id.New()

    trans := ut.New(enLocale, enLocale, idLocale)

    if err := registerTranslations(validate, trans); err != nil {
        return nil, err
    }

    return &structValidator{
        validate: validate,
        trans: trans,
    }, nil
}

func registerTranslations(validate *validator.Validate, uni *ut.UniversalTranslator) error {
    enTrans, _ := uni.GetTranslator("en")
    if err := en_trans.RegisterDefaultTranslations(validate, enTrans); err != nil {
        return fmt.Errorf("failed to register en translations: %w", err)
    }

    idTrans, _ := uni.GetTranslator("id")
    if err := id_trans.RegisterDefaultTranslations(validate, idTrans); err != nil {
        return fmt.Errorf("failed to register id translations: %w", err)
    }

    return nil
}
