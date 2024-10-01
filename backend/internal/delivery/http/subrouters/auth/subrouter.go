package authSubrouter

import (
	"github.com/go-chi/chi"
	"github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters"
)

type AuthSubrouter struct {
	subrouters.SubrouterDeps
}

const (
	accessJwtCookieName  = "accessJwt"
	refreshJwtCookieName = "refreshJwt"
)

func New(deps subrouters.SubrouterDeps) func(r chi.Router) {
	as := AuthSubrouter{deps}

	return func(r chi.Router) {
		r.Post("/register", as.Register)

		r.Get("/user", as.GetUser)
		r.Post("/user", as.UpdateUser)
		r.Delete("/user", as.DeleteUser)

		r.Post("/login", as.Login)
		r.Post("/logout", as.Logout)
		r.Post("/refresh", as.RefreshJwt)
		r.Post("/verify/jwt", as.VerifyJwt)

		r.Post("/verify/email/{token}", as.VerifyEmail)
	}
}
