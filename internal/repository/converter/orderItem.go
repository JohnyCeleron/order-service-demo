package converter

import (
	"order-service/internal/domain"
	"order-service/internal/repository/model"
)

func OrderItemDomainToModelDB(item domain.OrderItem, orderID string) model.OrderItem {
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

func OrderItemModelDBToDomain(item model.OrderItem) domain.OrderItem {
	return domain.OrderItem{
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
