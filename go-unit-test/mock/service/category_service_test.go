package service

import (
	"go-unit-test/mock/entity"
	"go-unit-test/mock/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock
var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryServiceGetNotFound(t *testing.T) {
	// program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	// call func
	result, err := categoryService.Get("1")
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestCategoryServiceGetSuccess(t *testing.T) {
	category := entity.Category{
		Id: "2",
		Name: "Laptop",
	}
	categoryRepository.Mock.On("FindById", category.Id).Return(category)

	result, err := categoryService.Get(category.Id)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}



