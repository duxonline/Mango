package catalogue

import "github.com/frankwang0/MangoCommerce/db"

type CategoryPgAdapter struct {
}

func (adapter *CategoryPgAdapter) Create(category Category) error {
	db := db.OpenDB()
	defer db.Close()

	result := db.Debug().Create(category)
	return result.Error
}
