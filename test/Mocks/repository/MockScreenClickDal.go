package repository

import (
	"StorageWorkerService/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockScreenClickDal struct {
	mock.Mock
}

func (m *MockScreenClickDal)Add(data *model.ScreenClickModel) error {
	args := m.Called(data)
	return  args.Error(0)
}
