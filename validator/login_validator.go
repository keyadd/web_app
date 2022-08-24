package validator

import (
	"github.com/go-playground/validator/v10"
	"web_app/model/request"
)

func LoginV(v *validator.Validate) {
	v.RegisterValidation("timing", request.LoginValidate)
}
