package repository

import (
	"StorageWorkerService/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockScreenSwipeDal struct {
	mock.Mock
}

func (m *MockScreenSwipeDal)Add(data *model.ScreenSwipeModel) error {
	args := m.Called(data)
	return  args.Error(0)
}
