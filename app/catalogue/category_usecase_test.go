package catalogue

import (
	"testing"

	"github.com/frankwang0/MangoCommerce/common"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("../../config/.env")
}

func Test_store_admins_add_parent_category(t *testing.T) {
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
	AssertCategory(t, category, createErr, name, 0)
}

func Test_store_admins_add_child_category(t *testing.T) {
	// Arrange
	var parentName = "Books"
	parentRequest := CreateCategoryRequest{
		Name:     parentName,
		ParentID: 0,
	}

	sut := CreateUseCase()
	parentCategory, _ := sut.CreateCategory(parentRequest)

	var categoryName = "Programming Books"
	childRequest := CreateCategoryRequest{
		Name:     categoryName,
		ParentID: parentCategory.ID,
	}

	// Act
	childCategory, createErr := sut.CreateCategory(childRequest)

	// Assert
	AssertCategory(t, childCategory, createErr, categoryName, parentCategory.ID)
}

func Test_retrieve_a_category(t *testing.T) {
	var name = "Books"
	var parentID = 0

	request := CreateCategoryRequest{
		Name:     name,
		ParentID: parentID,
	}
	sut := CreateUseCase()

	categoryCreated, _ := sut.CreateCategory(request)

	categoryFetched, fetchErr := sut.GetByID(categoryCreated.ID)

	AssertCategory(t, categoryFetched, fetchErr, name, 0)
}

func AssertCategory(t *testing.T, category *Category, appErr *common.AppError, name string, parentID int) {
	if appErr != nil {
		t.Errorf("An error has been thrown.")
	}

	if category == nil {
		t.Errorf("Current category is nil.")
	}

	if category.Name != name {
		t.Errorf("Category name is not saved correctly.")
	}

	if category.ParentID != parentID {
		t.Errorf("Category parentId is not saved correctly.")
	}
}
