package domain

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var orderValidator = validator.New()

func (o Order) Validate() (bool, error) {
	if err := orderValidator.Struct(o); err != nil {
		return false, fmt.Errorf("ошибка при валидации: %v", err)
	}
	if valid, err := o.Delivery.Validate(); !valid {
		return false, fmt.Errorf("ошибка при валидации delivery: %v", err)
	}
	if valid, err := o.Payment.Validate(); !valid {
		return false, fmt.Errorf("ошибка при валидации payment: %v", err)
	}
	for _, item := range o.Items {
		if valid, err := item.Validate(); !valid {
			return false, fmt.Errorf("ошибка при валидации item: %v", err)
		}
	}
	return true, nil
}
