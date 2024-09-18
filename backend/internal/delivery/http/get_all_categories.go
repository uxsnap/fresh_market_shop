package deliveryHttp

import (
	"net/http"
)

func (s *Server) getAllCategories(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}
