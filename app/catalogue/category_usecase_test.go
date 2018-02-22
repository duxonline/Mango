package catalogue

import (
	"testing"
)

func Test_store_admins_add_top_category(t *testing.T) {
	// Arrange
	var name = "Books"
	request := CreateCategoryRequest{
		Name:     name,
		ParentID: 0,
	}
	sut := CreateUseCase()

	// Act
	category, _ := sut.CreateCategory(request)

	// Assert
	if category.Name != name {
		t.Errorf("Category name is not saved correctly.")
	}

	if category.ParentID != 0 {
		t.Errorf("Category parentId is not saved correctly.")
	}
}
