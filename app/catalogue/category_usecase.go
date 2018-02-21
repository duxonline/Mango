package catalogue

import (
	"github.com/frankwang0/MangoCommerce/common"
)

type UseCase struct {
	CategoryRepo CategoryRepo
}

func (uc *UseCase) CreateCategory(request CreateCategoryRequest) (*Category, *common.AppError) {
	errCode := request.Validate()
	if errCode != common.AllGood {
		return nil, &common.AppError{ErrorCode: errCode, HttpStatusCode: 400}
	}
	return &Category{}, nil
}
