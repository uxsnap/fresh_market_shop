package deliveryHttp

import (
	"net/http"
	"time"
)

func (h *Handler) getHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("http server is working " + time.Now().String()))
}