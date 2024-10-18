package httpUtils

import (
	"encoding/json"
	"log"
	"net/http"
)

func DecodeJsonRequest(r *http.Request, dest interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		log.Printf("failed to decode request: %v", err)
		return err
	}
	return nil
}
