package repository

import (
	"StorageWorkerService/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockInventoryDal struct {
	mock.Mock
}

func (m *MockInventoryDal)Add(data *model.InventoryModel) error {
	args := m.Called(data)
	return  args.Error(0)
}
