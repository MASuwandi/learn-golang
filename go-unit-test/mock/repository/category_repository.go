package repository

import "go-unit-test/mock/entity"

// Contract
type CategoryRepository interface {
	// funcName(param) return
	FindById(id string) *entity.Category
}