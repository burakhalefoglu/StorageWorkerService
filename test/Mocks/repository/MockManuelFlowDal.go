package repository

import (
	"StorageWorkerService/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockManuelFlowDal struct {
	mock.Mock
}

func (m *MockManuelFlowDal)Add(data *model.ManuelFlowModel) error {
	args := m.Called(data)
	return  args.Error(0)
}
