package httpUtils

import "net/http"

func NewCookie(key string, value string) *http.Cookie {
	return &http.Cookie{
		Name:     key,
		Value:    value,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
}
