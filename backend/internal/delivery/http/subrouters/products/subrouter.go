package productsSubrouter

import (
	"github.com/go-chi/chi"
	"github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters"
)

type ProductsSubrouter struct {
	subrouters.SubrouterDeps
}

func New(deps subrouters.SubrouterDeps) func(r chi.Router) {
	ps := ProductsSubrouter{deps}

	return func(r chi.Router) {

		r.Get("/", ps.GetProducts)
		r.Get("/{uid}", ps.GetProduct)
		r.Get("/by_names", ps.getProductsByNames)
		r.Get("/new", ps.getNewProducts)

		r.Get("/count/{uid}", ps.GetProductCount)

		r.Post("/count", ps.SetProductCount)
		r.Post("/count/inc", ps.IncrementProductCount)
		r.Post("/count/dec", ps.DecrementProductCount)

		r.Post("/", ps.CreateProduct)
		r.Put("/", ps.UpdateProduct)
		r.Delete("/{uid}", ps.DeleteProduct)
	}
}
