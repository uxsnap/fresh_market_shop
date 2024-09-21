package deliveryHttp

import "github.com/go-chi/chi"

func (s *Server) CategoriesSubrouter(r chi.Router) {
	r.Get("/", s.getAllCategories)
}
