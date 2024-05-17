package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
	"time"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	_ = validate.RegisterValidation("phone_number", func(fl validator.FieldLevel) bool {
		phone := fl.Field().String()
		return strings.HasPrefix(phone, "+62") && len(phone) >= 12 && len(phone) <= 15
	})
	_ = validate.RegisterValidation("uppercase", func(fl validator.FieldLevel) bool {
		return regexp.MustCompile(`[A-Z]`).MatchString(fl.Field().String())
	})
	_ = validate.RegisterValidation("digit", func(fl validator.FieldLevel) bool {
		return regexp.MustCompile(`[0-9]`).MatchString(fl.Field().String())
	})
	_ = validate.RegisterValidation("special", func(fl validator.FieldLevel) bool {
		return regexp.MustCompile(`[^A-Za-z0-9]`).MatchString(fl.Field().String())
	})
	_ = validate.RegisterValidation("dob", func(fl validator.FieldLevel) bool {
		dateStr := fl.Field().String()
		_, err := time.Parse("2006-01-02", dateStr)
		return err == nil
	})
}

func Validate(i interface{}) error {
	return validate.Struct(i)
}
