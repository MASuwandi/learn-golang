package repository

import (
	"go-unit-test/mock/entity"
	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}


func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	}

	// conversion -> .(entity.Category)
	category := arguments.Get(0).(entity.Category)
	return &category
}