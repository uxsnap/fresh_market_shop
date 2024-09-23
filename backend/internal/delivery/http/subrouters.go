package deliveryHttp

import "github.com/go-chi/chi"

func (h *Handler) CategoriesSubrouter(r chi.Router) {
	r.Get("/", h.getAllCategories)
}
