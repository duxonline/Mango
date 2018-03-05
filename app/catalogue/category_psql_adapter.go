package catalogue

import "github.com/frankwang0/MangoCommerce/db"

type CategoryPsqlAdapter struct {
}

func (adapter *CategoryPsqlAdapter) Create(category *Category) error {
	db := db.OpenDB()
	defer db.Close()

	result := db.Debug().Create(category)
	return result.Error
}

func (adapter *CategoryPsqlAdapter) GetByID(categoryID int) (*Category, error) {
	db := db.OpenDB()
	defer db.Close()

	var category Category
	result := db.Debug().Where(&Category{ID: categoryID}).First(&category)
	if result.RecordNotFound() {
		return nil, nil
	}

	return &category, result.Error
}
