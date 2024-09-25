package deliveryHttp

import "github.com/go-chi/chi"

func (h *Handler) CategoriesSubrouter(r chi.Router) {
	r.Get("/", h.getAllCategories)
}

func (h *Handler) HealthSubrouter(r chi.Router) {
	r.Get("/", h.getHealthCheck)
}

func (h *Handler) AuthSubrouter(r chi.Router) {
	r.Post("/register", h.Register)

	r.Get("/user", h.GetUser)
	r.Post("/user", h.UpdateUser)
	r.Delete("/user", h.DeleteUser)

	r.Post("/login", h.Login)
	r.Post("/logout", h.Logout)
	r.Post("/refresh", h.RefreshJwt)
	r.Post("/verify/jwt", h.VerifyJwt)

	r.Post("/verify/email/{token}", h.VerifyEmail)
}
