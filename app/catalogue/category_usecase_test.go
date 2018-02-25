package catalogue

import (
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("../../config/.env")
}

func Test_store_admins_add_top_category(t *testing.T) {
	// Arrange
	var name = "Books"
	request := CreateCategoryRequest{
		Name:     name,
		ParentID: 0,
	}
	sut := CreateUseCase()

	// Act
	category, createErr := sut.CreateCategory(request)

	// Assert
	if createErr != nil {
		t.Errorf("Creating category throws an error.")
	}

	if category.Name != name {
		t.Errorf("Category name is not saved correctly.")
	}

	if category.ParentID != 0 {
		t.Errorf("Category parentId is not saved correctly.")
	}
}
