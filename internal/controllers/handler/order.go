package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	o "order-service/internal/service/order"
)

func (h *HttpHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "order_uid")
	if uid == "" {
		http.Error(w, `{"error":"order_uid is required"}`, http.StatusBadRequest)
		return
	}
	order, err := h.orderService.GetById(r.Context(), uid)
	if err != nil {
		if errors.Is(err, o.ErrOrderNotFound) {
			http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, `{"error":"internal"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(order)
}
