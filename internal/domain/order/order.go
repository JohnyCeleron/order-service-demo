package domain

import (
	"time"

	domainDelivery "order-service/internal/domain/delivery"
	domainOrderItem "order-service/internal/domain/orderItem"
	domainPayment "order-service/internal/domain/payment"
)

type Order struct {
	OrderUID          string                      `json:"order_uid" validate:"required"`
	TrackNumber       string                      `json:"track_number" validate:"required"`
	Entry             string                      `json:"entry" validate:"required"`
	Delivery          domainDelivery.Delivery     `json:"delivery" validate:"required"`
	Payment           domainPayment.Payment       `json:"payment" validate:"required"`
	Items             []domainOrderItem.OrderItem `json:"items" validate:"required"`
	Locale            string                      `json:"locale" validate:"required"`
	InternalSignature string                      `json:"internal_signature"`
	CustomerID        string                      `json:"customer_id" validate:"required"`
	DeliveryService   string                      `json:"delivery_service" validate:"required"`
	ShardKey          string                      `json:"shardkey" validate:"required"`
	SmID              int                         `json:"sm_id" validate:"required"`
	DateCreated       time.Time                   `json:"date_created" validate:"required"`
	OofShard          string                      `json:"oof_shard" validate:"required"`
}
