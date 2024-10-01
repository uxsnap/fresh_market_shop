package healthSubrouter

import (
	"net/http"
	"time"
)

func (h *HealthSubrouter) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("http server is working " + time.Now().String()))
}
