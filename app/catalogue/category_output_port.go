package catalogue

type CategoryRepo interface {
	Create(category *Category) error
	GetByID(categoryID int) (*Category, error)
}
