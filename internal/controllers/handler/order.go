package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"order-service/internal/controllers/response"
	o "order-service/internal/service/order"
)

// GetByUID godoc
// @Summary     Получить заказ по UID
// @Tags        order
// @Param       order_uid path string true "UID заказа" example(b563feb7b2b84b6test)
// @Produce     json
// @Success     200 {object} response.OrderResponse
// @Failure     404 {object} response.ErrorResponse
// @Failure     500 {object} response.ErrorResponse
// @Router      /order/{order_uid} [get]
func (h *HttpHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "order_uid")
	w.Header().Set("Content-Type", "application/json")
	if uid == "" {
		writeJSON(w, http.StatusBadRequest, response.NewErrorResponse("order_uid is required"))
		return
	}

	order, err := h.orderService.GetById(r.Context(), uid)
	if err != nil {
		if errors.Is(err, o.ErrOrderNotFound) {
			writeJSON(w, http.StatusNotFound,
				response.NewErrorResponse(fmt.Sprintf("order with id %v not found", uid)))
			return
		}
		writeJSON(w, http.StatusInternalServerError,
			response.NewErrorResponse("internal"))
		return
	}

	writeJSON(w, http.StatusOK, response.OrderResponse(order))
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, `{"error":"encode error"}`, http.StatusInternalServerError)
	}
}
