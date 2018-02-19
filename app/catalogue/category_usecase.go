package catalogue

type UseCase struct {
	// BcRequestRepo    BcRequestRepo
}

func (uc *UseCase) AddCategory(request AddCategoryRequest) (*Category, error) {
	return &Category{}, nil
}
