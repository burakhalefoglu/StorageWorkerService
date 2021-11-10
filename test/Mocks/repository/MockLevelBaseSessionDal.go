package repository

import (
	"StorageWorkerService/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockLevelBaseSessionDal struct {
	mock.Mock
}

func (m *MockLevelBaseSessionDal)Add(data *model.LevelBaseSessionModel) error {
	args := m.Called(data)
	return  args.Error(0)
}
