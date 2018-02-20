package catalogue

import (
	"github.com/frankwang0/MangoCommerce/common"
)

type UseCase struct {
	CategoryRepo CategoryRepo
}

func (uc *UseCase) CreateCategory(request CreateCategoryRequest) (*Category, *common.AppError) {
	return &Category{}, nil
}
