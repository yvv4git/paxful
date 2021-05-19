package usecases

import "github.com/asaskevich/govalidator"

// Validate is used for validation all form in one point.
func Validate(form interface{}) (err error) {
	result, err := govalidator.ValidateStruct(form)
	if err == nil && !result {
		return ErrorUnknown
	}
	return
}
