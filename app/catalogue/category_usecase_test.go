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
	request := CreateCategoryRequest{
		Name:     name,
		ParentID: 0,
	}
	sut := CreateUseCase()

	categoryCreated, _ := sut.CreateCategory(request)

	categoryFetched, fetchErr := sut.GetByID(categoryCreated.id)

	if fetchErr != nil {
		t.Errorf("Fetching category throws an error.")
	}

	if categoryFetched == nil {
		t.Errorf("Fetching category return nil.")
	}

	if categoryFetched.Name != name {
		t.Errorf("Category name is not correct.")
	}
}

func AssertCategory(t *testing.T, category *Category, createErr *common.AppError, name string, parentID int) {
	if createErr != nil {
		t.Errorf("Creating category throws an error.")
	}

	if category.Name != name {
		t.Errorf("Category name is not saved correctly.")
	}

	if category.ParentID != parentID {
		t.Errorf("Category parentId is not saved correctly.")
	}
}
