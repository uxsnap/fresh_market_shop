package subrouters

import (
	"context"
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

func (m *Middleware) verifyToken(ctx context.Context, tokenCookie string) (*httpEntity.AuthUserInfo, *errorWrapper.Error) {
	if err := m.VerifyJwt(ctx, tokenCookie); err != nil {
		log.Printf("auth middleware: invalid jwt token: %v", err)
		return nil, errorWrapper.NewError(errorWrapper.JwtAuthMiddleware, "неправильный токен")
	}

	token, _ := jwt.Parse(tokenCookie, nil)
	if token == nil {
		log.Printf("auth middleware: failed to parse jwt token: empty token")
		return nil, errorWrapper.NewError(errorWrapper.JwtAuthMiddleware, "неправильный токен")
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	userUidStr, ok := claims["user_uid"]
	if !ok {
		log.Printf("auth middleware: invalid jwt: user_uid is empty")
		return nil, errorWrapper.NewError(errorWrapper.JwtAuthMiddleware, "неправильный токен")
	}

	userUid, err := uuid.FromString(userUidStr.(string))
	if err != nil {
		log.Printf("auth middleware: invalid jwt: user_uid is invalid")
		return nil, errorWrapper.NewError(errorWrapper.JwtAuthMiddleware, "неправильный токен")
	}

	userRole, ok := claims["role"]
	if !ok {
		log.Printf("auth middleware: invalid jwt: role is empty")
		return nil, errorWrapper.NewError(errorWrapper.JwtAuthMiddleware, "неправильный токен")
	}

	userPermissions, ok := claims["permissions"]
	if !ok {
		log.Printf("auth middleware: invalid jwt: permissions is empty")
		return nil, errorWrapper.NewError(errorWrapper.JwtAuthMiddleware, "неправильный токен")
	}

	userInfo := httpEntity.AuthUserInfo{
		UserUid:     userUid,
		Role:        userRole.(string),
		Permissions: userPermissions.(string),
	}

	return &userInfo, nil
}

func (m *Middleware) AuthOrPass(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tokenCookie := httpUtils.GetBearerToken(r)

		if tokenCookie == "" {
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		userInfo, errorWrapper := m.verifyToken(ctx, tokenCookie)

		if errorWrapper != nil {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper)
			return
		}

		ctx = httpEntity.ContextWithAuthUserInfo(ctx, *userInfo)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

func (m *Middleware) Auth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tokenCookie := httpUtils.GetBearerToken(r)

		if tokenCookie == "" {
			log.Printf("отсутствует токен")
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(errorWrapper.JwtAuthMiddleware, "отсутствует токен"))
			return
		}

		userInfo, errorWrapper := m.verifyToken(ctx, tokenCookie)

		if errorWrapper != nil {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper)
			return
		}

		ctx = httpEntity.ContextWithAuthUserInfo(ctx, *userInfo)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
