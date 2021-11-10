package repository

import (
	"StorageWorkerService/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockOfferBehaviorDal struct {
	mock.Mock
}

func (m *MockOfferBehaviorDal)Add(data *model.OfferBehaviorModel) error {
	args := m.Called(data)
	return  args.Error(0)
}
