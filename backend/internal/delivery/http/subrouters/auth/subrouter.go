package authSubrouter

import (
	"github.com/go-chi/chi"
	"github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters"
)

type AuthSubrouter struct {
	subrouters.SubrouterDeps
}

func New(deps subrouters.SubrouterDeps) func(r chi.Router) {
	as := AuthSubrouter{deps}

	return func(r chi.Router) {
		r.Post("/register", as.Register)

		r.Get("/user", as.GetAuthUser)
		r.Post("/user", as.UpdateAuthUser)
		r.Delete("/user", as.DeleteAuthUser)

		r.Post("/login", as.Login)
		r.Post("/logout", as.Logout)
		r.Post("/refresh", as.RefreshJwt)
		r.Post("/verify/jwt", as.VerifyJwt)

		r.Post("/verify/email/{token}", as.VerifyEmail)

		r.Group(func(r chi.Router) {
			r.Use(as.Middleware.AuthOrPass)

			r.Post("/admin", as.CreateAdmin)
			r.Get("/admins", as.GetAdmins)
		})
	}
}
