package catalogue

type InputPort interface {
	AddCategory(request AddCategoryRequest) (*Category, error)
}
