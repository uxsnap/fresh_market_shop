package assetsSubrouter

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
	formUtils "github.com/uxsnap/fresh_market_shop/backend/internal/utils"
)

func handleIsFileMatches(curFilePath string) bool {
	matches, err := filepath.Glob(curFilePath)

	if len(matches) == 0 || err != nil {
		return false
	}

	if !formUtils.IsImageExtensionAllowed(matches[0]) {
		return false
	}

	return true
}

func (as *AssetsSubrouter) getUserPhoto(path string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		uid, err := uuid.FromString(chi.URLParam(r, "user_uid"))
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(
				errorWrapper.UserInfoError, "не удалось получить информацию о пользователе",
			))
			return
		}

		userInfo, err := httpEntity.AuthUserInfoFromContext(r.Context())
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(
				errorWrapper.UserInfoError, "не удалось получить информацию о пользователе",
			))
			return
		}

		if userInfo.UserUid != uid {
			httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(
				errorWrapper.UserInfoError, "не совпадают идентификаторы",
			))
			return
		}

		workDir, _ := os.Getwd()

		curFilePath := filepath.Join(filepath.Join(workDir, path), uid.String()) + ".webp"

		ok := handleIsFileMatches(curFilePath)

		if !ok {
			httpUtils.WriteErrorResponse(w, http.StatusNotFound, nil)
			return
		}

		pathPrefix := strings.TrimSuffix(chi.RouteContext(r.Context()).RoutePattern(), "/{user_uid}.webp")

		fs := http.StripPrefix(pathPrefix, http.FileServer(http.Dir(path)))

		fs.ServeHTTP(w, r)
	}
}
