package usersSubrouter

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/consts"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
	formUtils "github.com/uxsnap/fresh_market_shop/backend/internal/utils"
)

const MAX_PHOTO_SIZE = 15

func (h *UsersSubrouter) uploadPhoto(w http.ResponseWriter, r *http.Request) {
	uid, err := uuid.FromString(chi.URLParam(r, "user_uid"))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(errorWrapper.UserInfoError, "user_uid не найден"))
		return
	}

	r.ParseMultipartForm(5 << 20)

	file, handler, err := r.FormFile("photo")
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(
			errorWrapper.UserPhotoError, "не удалось загрузить фото",
		))
		return
	}

	if !formUtils.IsImageExtensionAllowed(handler.Filename) {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(
			errorWrapper.UserPhotoError, "неподдерживаемый формат",
		))
		return
	}

	defer file.Close()

	uploadPath, err := os.Getwd()

	if err != nil {
		log.Println(err)
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(
			errorWrapper.UserPhotoError, "не удалось загрузить фото на сервер",
		))
		return
	}

	newFileName := fmt.Sprintf("%v.%v", uid.String(), formUtils.GetFileExtension(handler.Filename))

	dst, err := os.Create(filepath.Join(
		uploadPath, consts.USER_PHOTO_PATH, newFileName,
	))

	if err != nil {
		log.Println(err)
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(
			errorWrapper.UserPhotoError, "не удалось загрузить фото на сервер",
		))
		return
	}

	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		log.Println(err)
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(
			errorWrapper.UserPhotoError, "не удалось загрузить фото на сервер",
		))
		return
	}

	w.WriteHeader(http.StatusOK)
}
