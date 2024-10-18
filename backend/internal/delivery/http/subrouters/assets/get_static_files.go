package assetsSubrouter

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
)

const PATH = "assets/imgs"

func (as *AssetsSubrouter) getStaticFiles(w http.ResponseWriter, r *http.Request) {
	workDir, _ := os.Getwd()

	filesDir := http.Dir(filepath.Join(workDir, PATH))

	routeCtx := chi.RouteContext(r.Context())

	pathPrefix := strings.TrimSuffix(routeCtx.RoutePattern(), "/*")

	fs := http.StripPrefix(pathPrefix, http.FileServer(filesDir))

	fs.ServeHTTP(w, r)
}
