package converter

import (
	domainOrder "order-service/internal/domain/order"
	domainOrderItem "order-service/internal/domain/orderItem"
	"order-service/internal/repository/model"
)

func OrderModelDBToDomain(dbOrder model.Order) domainOrder.Order {
	domainItems := make([]domainOrderItem.OrderItem, len(dbOrder.Items))
	for i, item := range dbOrder.Items {
		domainItems[i] = OrderItemModelDBToDomain(item)
	}
	return domainOrder.Order{
		OrderUID:          dbOrder.ID,
		TrackNumber:       dbOrder.TrackNumber,
		Entry:             dbOrder.Entry,
		Delivery:          DeliveryModelDBToDomain(dbOrder.Delivery),
		Payment:           PaymentModelDBToDomain(dbOrder.Payment),
		Items:             domainItems,
		Locale:            dbOrder.Locale,
		InternalSignature: dbOrder.InternalSignature,
		CustomerID:        dbOrder.CustomerID,
		DeliveryService:   dbOrder.DeliveryService,
		ShardKey:          dbOrder.Shardkey,
		SmID:              dbOrder.SmID,
		DateCreated:       dbOrder.DateCreated,
		OofShard:          dbOrder.OofShard,
	}
}

func OrderDomainToModelDB(domainOrder domainOrder.Order) model.Order {
	modelItems := make([]model.OrderItem, len(domainOrder.Items))
	for i, item := range domainOrder.Items {
		modelItems[i] = OrderItemDomainToModelDB(item, domainOrder.OrderUID)
	}
	return model.Order{
		ID:                domainOrder.OrderUID,
		TrackNumber:       domainOrder.TrackNumber,
		Entry:             domainOrder.Entry,
		Locale:            domainOrder.Locale,
		InternalSignature: domainOrder.InternalSignature,
		CustomerID:        domainOrder.CustomerID,
		DeliveryService:   domainOrder.DeliveryService,
		Shardkey:          domainOrder.ShardKey,
		SmID:              domainOrder.SmID,
		DateCreated:       domainOrder.DateCreated,
		OofShard:          domainOrder.OofShard,
		Delivery:          DeliveryDomainToModelDB(domainOrder.Delivery),
		Payment:           PaymentDomainToModelDB(domainOrder.Payment),
		Items:             modelItems,
	}
}
