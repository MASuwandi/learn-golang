package service

import (
	"errors"
	"go-unit-test/mock/entity"
	"go-unit-test/mock/repository"
)

type CategoryService struct {
	// attribute	type
	Repository		repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return category, errors.New("category not found")
	}
	return category, nil
}