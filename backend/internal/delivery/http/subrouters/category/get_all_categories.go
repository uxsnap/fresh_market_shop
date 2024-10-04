package categorySubrouter

import (
	"context"
	"net/http"
)

func (h *CategorySubrouter) getAllCategories(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	h.CategoriesService.GetAllCategories(ctx)

	w.Write([]byte("test"))
}