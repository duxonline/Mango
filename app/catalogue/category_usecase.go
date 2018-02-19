package catalogue

type UseCase struct {
	CategoryRepo CategoryRepo
}

func (uc *UseCase) AddCategory(request AddCategoryRequest) (*Category, error) {
	return &Category{}, nil
}
