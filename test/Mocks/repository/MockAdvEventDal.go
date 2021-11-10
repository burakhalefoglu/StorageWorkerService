package repository

import (
	"StorageWorkerService/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockAdvEventDal struct {
	mock.Mock
}

func (m *MockAdvEventDal) Add(data *model.AdvEventDataModel) error {
	args := m.Called(data)
	return  args.Error(0)
}