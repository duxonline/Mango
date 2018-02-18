package catalogue

import "net/http"

func AddCategory(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return nil, nil
}

func GetCategory(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return "this is a test category", nil
}
