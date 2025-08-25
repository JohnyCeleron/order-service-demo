package handler

import "order-service/internal/service/order"

type HttpHandler struct {
	orderService *order.OrderService
}

func New(orderService *order.OrderService) *HttpHandler {
	return &HttpHandler{
		orderService: orderService,
	}
}
