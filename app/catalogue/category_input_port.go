package catalogue

import (
	"github.com/frankwang0/MangoCommerce/common"
)

type InputPort interface {
	CreateCategory(request CreateCategoryRequest) (*Category, *common.AppError)
}
