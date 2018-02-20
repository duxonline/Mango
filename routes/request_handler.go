package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/frankwang0/MangoCommerce/common"
)

type RequestHandler func(w http.ResponseWriter, r *http.Request) (interface{}, *common.AppError)

func (fn RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	result, appErr := fn(w, r)
	if appErr != nil {
		fmt.Println(appErr)
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
