package domain

import (
	"time"

	domainDelivery "order-service/internal/domain/delivery"
	domainOrderItem "order-service/internal/domain/orderItem"
	domainPayment "order-service/internal/domain/payment"
)

type Order struct {
	OrderUID          string                      `json:"order_uid" validate:"required" example:"b563feb7b2b84b6test"`
	TrackNumber       string                      `json:"track_number" validate:"required" example:"WBILMTESTTRACK"`
	Entry             string                      `json:"entry" validate:"required" example:"WBIL"`
	Delivery          domainDelivery.Delivery     `json:"delivery" validate:"required"`
	Payment           domainPayment.Payment       `json:"payment" validate:"required"`
	Items             []domainOrderItem.OrderItem `json:"items" validate:"required"`
	Locale            string                      `json:"locale" validate:"required" example:"en"`
	InternalSignature string                      `json:"internal_signature" example:""`
	CustomerID        string                      `json:"customer_id" validate:"required" example:"test"`
	DeliveryService   string                      `json:"delivery_service" validate:"required" example:"meest"`
	ShardKey          string                      `json:"shardkey" validate:"required" example:"9"`
	SmID              int                         `json:"sm_id" validate:"required" example:"99"`
	DateCreated       time.Time                   `json:"date_created" validate:"required" example:"2021-11-26T06:22:19Z"`
	OofShard          string                      `json:"oof_shard" validate:"required" example:"1"`
}
