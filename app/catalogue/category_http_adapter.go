package catalogue

import (
	"encoding/json"
	"net/http"
)

func AddCategory(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	request := new(AddCategoryRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, err
	}

	useCase := CreateUseCase()
	return useCase.AddCategory(*request)
}

func GetCategory(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return "this is a test category", nil
}
