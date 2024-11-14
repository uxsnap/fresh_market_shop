package assetsSubrouter

import (
	"github.com/go-chi/chi"
	"github.com/uxsnap/fresh_market_shop/backend/internal/consts"
	"github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters"
)

type AssetsSubrouter struct {
	subrouters.SubrouterDeps
}

func New(deps subrouters.SubrouterDeps) func(r chi.Router) {
	as := AssetsSubrouter{deps}

	return func(r chi.Router) {
		r.Get("/imgs/*", as.getStaticFiles("assets/imgs"))
		r.Get("/recipes/*", as.getStaticFiles("assets/recipes"))

		r.Group(func(r chi.Router) {
			r.Use(as.Middleware.Auth)

			r.Get("/photo/{user_uid}.webp", as.getUserPhoto(consts.USER_PHOTO_PATH))
		})
	}
}
