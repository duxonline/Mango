package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type RequestHandler func(w http.ResponseWriter, r *http.Request) (interface{}, error)

func (fn RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	result, err := fn(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}

	writeResponse(w, r, result)
}

func writeResponse(w http.ResponseWriter, r *http.Request, result interface{}) {
	var statusCode int
	if strings.EqualFold(r.Method, "POST") {
		statusCode = http.StatusCreated
	} else {
		statusCode = http.StatusOK
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(result)
}
