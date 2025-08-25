package domain

type Payment struct {
	Transaction  string `json:"transaction" validate:"required" example:"b563feb7b2b84b6test"`
	RequestID    string `json:"request_id" example:""`
	Currency     string `json:"currency" validate:"required" example:"USD"`
	Provider     string `json:"provider" validate:"required" example:"wbpay"`
	Amount       int    `json:"amount" validate:"gt=0"  example:"1817"`
	PaymentDt    int    `json:"payment_dt" validate:"required" example:"1637907727"`
	Bank         string `json:"bank" validate:"required" example:"alpha"`
	DeliveryCost int    `json:"delivery_cost" validate:"gte=0" example:"1500"`
	GoodsTotal   int    `json:"goods_total" validate:"gt=0" example:"317"`
	CustomFee    int    `json:"custom_fee" validate:"gte=0" example:"0"`
}
