package categoriesSubrouter

import (
	"github.com/go-chi/chi"
	"github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters"
)

type CategoriesSubrouter struct {
	subrouters.SubrouterDeps
}

func New(deps subrouters.SubrouterDeps) func(r chi.Router) {
	cs := CategoriesSubrouter{deps}

	return func(r chi.Router) {
		r.Get("/", cs.getAllCategories)
		r.Get("/{uid}", cs.getCategoryByUid)
		r.Get("/{category_uid}/products", cs.getCategoryProducts)

		r.Group(func(r chi.Router) {
			r.Use(cs.Middleware.Auth)

			r.Post("/", cs.createCategory)
			r.Put("/", cs.updateCategory)
			r.Delete("/", cs.deleteCategory)
		})
	}
}
