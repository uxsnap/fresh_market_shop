package httpUtils

import (
	"net/http"
	"strings"
)

func GetBearerToken(r *http.Request) string {
	bearer := r.Header.Get("Authorization")

	return strings.Replace(bearer, "Bearer ", "", 1)
}
