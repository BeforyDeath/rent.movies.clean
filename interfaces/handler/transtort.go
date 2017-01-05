package handler

import (
	"encoding/json"
	"io"
	"net/http"
)

type response struct {
	w       http.ResponseWriter
	Success bool
	Data    interface{}
	Error   interface{}
}

func newResponse(w http.ResponseWriter) response {
	return response{w: w}
}

func (r *response) Send() {
	r.w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(r.w).Encode(r)
}

func jsonRequest(r *http.Request) (map[string]interface{}, error) {
	defer r.Body.Close()
	var JSON map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	decoder.UseNumber()
	err := decoder.Decode(&JSON)
	if err != nil {
		if err != io.EOF {
			return JSON, err
		}
	}
	return JSON, nil
}
