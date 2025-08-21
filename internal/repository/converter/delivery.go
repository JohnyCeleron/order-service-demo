package converter

import (
	domainDelivery "order-service/internal/domain/delivery"
	"order-service/internal/repository/model"
)

func DeliveryDomainToModelDB(d domainDelivery.Delivery) model.Delivery {
	return model.Delivery{
		Name:    d.Name,
		Phone:   d.Phone,
		Zip:     d.Zip,
		City:    d.City,
		Address: d.Address,
		Region:  d.Region,
		Email:   d.Email,
	}
}

func DeliveryModelDBToDomain(db model.Delivery) domainDelivery.Delivery {
	return domainDelivery.Delivery{
		Name:    db.Name,
		Phone:   db.Phone,
		Zip:     db.Zip,
		City:    db.City,
		Address: db.Address,
		Region:  db.Region,
		Email:   db.Email,
	}
}
