package deliveryHttp

import (
	"context"
	"net/http"
)

func (h *Handler) getAllCategories(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	h.categoriesService.GetAllCategories(ctx)

	w.Write([]byte("test"))
}
