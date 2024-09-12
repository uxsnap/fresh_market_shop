package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func initCategoryRoute(r chi.Router) {
	r.Get("/", getAllCategories)
}

func getAllCategories(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("test")))
}
