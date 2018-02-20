package catalogue

import (
	"encoding/json"
	"net/http"

	"github.com/frankwang0/MangoCommerce/common"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) (interface{}, *common.AppError) {
	request := new(CreateCategoryRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, &common.AppError{Error: err}
	}

	useCase := CreateUseCase()
	return useCase.CreateCategory(*request)
}

func GetCategory(w http.ResponseWriter, r *http.Request) (interface{}, *common.AppError) {
	return "this is a test category", nil
}
