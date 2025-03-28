package repository

import (
	"github.com/stretchr/testify/mock"
	"go-unit-test/entity"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	args := repository.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	}
	category := args.Get(0).(entity.Category)
	return &category
}
