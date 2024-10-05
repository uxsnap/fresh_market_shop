package ordersSubrouter

import "net/http"

func (h *OrdersSubrouter) getHistory(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
