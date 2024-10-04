package categorySubrouter

import (
	"github.com/go-chi/chi"
	"github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters"
)

type CategorySubrouter struct {
	subrouters.SubrouterDeps
}

func New(deps subrouters.SubrouterDeps) func(r chi.Router) {
	cs := CategorySubrouter{deps}

	return func(r chi.Router) {
		r.Get("/", cs.getAllCategories)
		r.Get("/{uid}", cs.getCategoryByUid)
		r.Get("/{category_uid}/products", cs.getCategoryProducts)

		r.Post("/manage/create", cs.createCategory)
		r.Post("/manage/update", cs.updateCategory)
		r.Post("/manage/delete", cs.deleteCategory)
	}
}
