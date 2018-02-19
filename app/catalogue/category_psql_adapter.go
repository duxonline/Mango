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
