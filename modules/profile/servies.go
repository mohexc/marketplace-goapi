package profile

import (
	"marketplace-goapi/modules/base"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(profileReq CreateProfileRequest) []*base.ErrorResponse {
	var errors []*base.ErrorResponse
	err := validate.Struct(profileReq)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element base.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func CreateProfile(profileReq *CreateProfileRequest) (*ProfileResponse, error) {
	return nil, nil
}
