package domain

import "github.com/go-playground/validator/v10"

var orderItemValidator = validator.New()

func (d OrderItem) Validate() (bool, error) {
	if err := orderItemValidator.Struct(d); err != nil {
		return false, err
	}
	return true, nil
}
