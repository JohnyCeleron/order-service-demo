package converter

import (
	domainPayment "order-service/internal/domain/payment"
	"order-service/internal/repository/model"
)

func PaymentDomainToModelDB(p domainPayment.Payment) model.Payment {
	return model.Payment{
		Transaction:  p.Transaction,
		RequestID:    p.RequestID,
		Currency:     p.Currency,
		Provider:     p.Provider,
		Amount:       p.Amount,
		PaymentDt:    p.PaymentDt,
		Bank:         p.Bank,
		DeliveryCost: p.DeliveryCost,
		GoodsTotal:   p.GoodsTotal,
		CustomFee:    p.CustomFee,
	}
}

func PaymentModelDBToDomain(db model.Payment) domainPayment.Payment {
	return domainPayment.Payment{
		Transaction:  db.Transaction,
		RequestID:    db.RequestID,
		Currency:     db.Currency,
		Provider:     db.Provider,
		Amount:       db.Amount,
		PaymentDt:    db.PaymentDt,
		Bank:         db.Bank,
		DeliveryCost: db.DeliveryCost,
		GoodsTotal:   db.GoodsTotal,
		CustomFee:    db.CustomFee,
	}
}
