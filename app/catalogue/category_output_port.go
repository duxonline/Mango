package catalogue

type CategoryRepo interface {
	Create(category *Category) error
}
