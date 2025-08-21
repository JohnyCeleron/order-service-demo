package domain

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var paymentValidator = validator.New()

func (p Payment) Validate() (bool, error) {
	if err := paymentValidator.Struct(p); err != nil {
		return false, err
	}
	if p.Amount != p.GoodsTotal+p.DeliveryCost+p.CustomFee {
		return false, fmt.Errorf("amount consistency error")
	}
	return true, nil
}
