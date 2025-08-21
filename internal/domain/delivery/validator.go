package domain

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	emailRegex        = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	phoneRegex        = regexp.MustCompile(`^\+?[0-9]{1,3}?[-.\s]?\(?[0-9]{1,4}?\)?[-.\s]?[0-9]{1,4}[-.\s]?[0-9]{1,9}$`)
	digitRegex        = regexp.MustCompile(`[0-9]`)
	deliveryValidator = validator.New()
)

func (d Delivery) Validate() (bool, error) {
	if err := deliveryValidator.Struct(d); err != nil {
		return false, err
	}
	if valid, err := d.validateEmail(); !valid {
		return false, err
	}
	if valid, err := d.validatePhone(); !valid {
		return false, err
	}
	return true, nil
}

func (d Delivery) validateEmail() (bool, error) {
	if !emailRegex.MatchString(d.Email) {
		return false, fmt.Errorf("email имеет невалидный формат")
	}
	return true, nil
}

func (d Delivery) validatePhone() (bool, error) {
	if !phoneRegex.MatchString(d.Phone) {
		return false, fmt.Errorf("номер телефона имеет невалидный формат")
	}
	digits := digitRegex.FindAllString(d.Phone, -1)
	if len(digits) < 10 {
		return false, fmt.Errorf("номер телефона должен иметь как минимум 10 цифр")
	}
	return true, nil
}
