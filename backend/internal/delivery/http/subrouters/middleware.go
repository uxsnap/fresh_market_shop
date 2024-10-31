package subrouters

import (
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

type Middleware struct {
	AuthService
}

func NewMiddleware(authService AuthService) *Middleware {
	return &Middleware{
		AuthService: authService,
	}
}

const accessJwtCookieName = "access_jwt"

func (m *Middleware) Auth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tokenCookie, err := r.Cookie(accessJwtCookieName)
		if err != nil {
			log.Printf("auth middleware: failed to get jwt token: %v", err)
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(err.Error(), "auth middleware: failed to get jwt token"))
			return
		}

		if err := m.VerifyJwt(ctx, tokenCookie.Value); err != nil {
			log.Printf("auth middleware: invalid jwt token: %v", err)
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(err.Error(), "auth middleware: invalid jwt token"))
			return
		}

		token, _ := jwt.Parse(tokenCookie.Value, nil)
		if token == nil {
			log.Printf("auth middleware: failed to parse jwt token: empty token")
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError("empty token", "auth middleware: failed to parse jwt token"))
			return
		}
		claims, _ := token.Claims.(jwt.MapClaims)
		userUidStr, ok := claims["user_uid"]
		if !ok {
			log.Printf("auth middleware: invalid jwt: user_uid is empty")
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError("", "auth middleware: invalid jwt"))
			return
		}

		userUid, err := uuid.FromString(userUidStr.(string))
		if err != nil {
			log.Printf("auth middleware: invalid jwt: user_uid is invalid")
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError("", "auth middleware: invalid jwt"))
			return
		}

		userRole, ok := claims["role"]
		if !ok {
			log.Printf("auth middleware: invalid jwt: role is empty")
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError("", "auth middleware: invalid jwt"))
			return
		}

		userPermissions, ok := claims["permissions"]
		if !ok {
			log.Printf("auth middleware: invalid jwt: permissions is empty")
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError("", "auth middleware: invalid jwt"))
			return
		}

		userInfo := httpEntity.AuthUserInfo{
			UserUid:     userUid,
			Role:        userRole.(string),
			Permissions: userPermissions.(string),
		}

		ctx = httpEntity.ContextWithAuthUserInfo(ctx, userInfo)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
