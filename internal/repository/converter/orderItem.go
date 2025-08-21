package converter

import (
	domainOrderItem "order-service/internal/domain/orderItem"
	"order-service/internal/repository/model"
)

func OrderItemDomainToModelDB(item domainOrderItem.OrderItem, orderID string) model.OrderItem {
	return model.OrderItem{
		ChrtID:      item.ChrtID,
		TrackNumber: item.TrackNumber,
		Price:       item.Price,
		Rid:         item.Rid,
		Name:        item.Name,
		Sale:        item.Sale,
		Size:        item.Size,
		TotalPrice:  item.TotalPrice,
		NmID:        item.NmID,
		Brand:       item.Brand,
		Status:      item.Status,
		OrderID:     orderID,
	}
}

func OrderItemModelDBToDomain(item model.OrderItem) domainOrderItem.OrderItem {
	return domainOrderItem.OrderItem{
		ChrtID:      item.ChrtID,
		TrackNumber: item.TrackNumber,
		Price:       item.Price,
		Rid:         item.Rid,
		Name:        item.Name,
		Sale:        item.Sale,
		Size:        item.Size,
		TotalPrice:  item.TotalPrice,
		NmID:        item.NmID,
		Brand:       item.Brand,
		Status:      item.Status,
	}
}
