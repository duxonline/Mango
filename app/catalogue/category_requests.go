package catalogue

import "github.com/frankwang0/MangoCommerce/common"

type CreateCategoryRequest struct {
	Name     string
	ParentID int
}

func (request CreateCategoryRequest) Validate() int {
	if request.Name == "" {
		return common.CategoryNameIsRequired
	}
	return common.AllGood
}
