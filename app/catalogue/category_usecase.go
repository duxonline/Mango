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

	category := newCategory(request)

	createErr := uc.CategoryRepo.Create(&category)
	if createErr != nil {
		return nil, &common.AppError{Error: createErr, HttpStatusCode: 500}
	}

	return &category, nil
}
