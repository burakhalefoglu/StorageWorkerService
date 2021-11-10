package repository

import (
	"StorageWorkerService/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockLocationDal struct {
	mock.Mock
}

func (m *MockLocationDal)Add(data *model.LocationModel) error {
	args := m.Called(data)
	return  args.Error(0)
}
