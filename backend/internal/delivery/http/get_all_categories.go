package deliveryHttp

import (
	"net/http"
)

func (h *Handler) getAllCategories(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}
