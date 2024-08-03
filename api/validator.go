package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type structValidator struct {
	validate *validator.Validate
}

func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

func createConfig() fiber.Config {
	return fiber.Config{
		StructValidator: &structValidator{validate: validator.New()},
	}
}
