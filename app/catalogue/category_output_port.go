package catalogue

type CategoryRepo interface {
	Create(request *AddCategoryRequest) error
}
