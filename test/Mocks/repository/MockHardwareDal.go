package repository

import (
	"StorageWorkerService/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockHardwareDal struct {
	mock.Mock
}

func (m *MockHardwareDal)Add(data *model.HardwareModel) error {
	args := m.Called(data)
	return  args.Error(0)
}

