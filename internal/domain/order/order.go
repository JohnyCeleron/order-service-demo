package domain

import (
	"time"

	domainDelivery "order-service/internal/domain/delivery"
	domainOrderItem "order-service/internal/domain/orderItem"
	domainPayment "order-service/internal/domain/payment"
)

type Order struct {
	OrderUID          string                      `json:"order_uid"`
	TrackNumber       string                      `json:"track_number"`
	Entry             string                      `json:"entry"`
	Delivery          domainDelivery.Delivery     `json:"delivery"`
	Payment           domainPayment.Payment       `json:"payment"`
	Items             []domainOrderItem.OrderItem `json:"items"`
	Locale            string                      `json:"locale"`
	InternalSignature string                      `json:"internal_signature"`
	CustomerID        string                      `json:"customer_id"`
	DeliveryService   string                      `json:"delivery_service"`
	ShardKey          string                      `json:"shardkey"`
	SmID              int                         `json:"sm_id"`
	DateCreated       time.Time                   `json:"date_created"`
	OofShard          string                      `json:"oof_shard"`
}
