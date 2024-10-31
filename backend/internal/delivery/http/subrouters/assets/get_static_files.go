package assetsSubrouter

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
)

func (as *AssetsSubrouter) getStaticFiles(path string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		workDir, _ := os.Getwd()

		workPath := path

		filesDir := http.Dir(filepath.Join(workDir, workPath))

		routeCtx := chi.RouteContext(r.Context())

		pathPrefix := strings.TrimSuffix(routeCtx.RoutePattern(), "/*")

		fmt.Println(pathPrefix)

		fs := http.StripPrefix(pathPrefix, http.FileServer(filesDir))

		fs.ServeHTTP(w, r)
	}
}
