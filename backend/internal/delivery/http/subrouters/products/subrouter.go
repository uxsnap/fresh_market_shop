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

		r.Get("/count/{uid}", ps.GetProductCount)

		r.Post("/manage/count", ps.SetProductCount)
		r.Post("/manage/count/inc", ps.IncrementProductCount)
		r.Post("/manage/count/dec", ps.DecrementProductCount)

		r.Post("/manage/create", ps.CreateProduct)
		r.Post("/manage/update", ps.UpdateProduct)
		r.Post("/manage/delete/{uid}", ps.DeleteProduct)
	}
}
