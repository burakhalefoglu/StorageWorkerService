package repository

import (
	"StorageWorkerService/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockEnemyBaseLoginLevelDal struct {
	mock.Mock
}

func (m *MockEnemyBaseLoginLevelDal)Add(data *model.EnemyBaseLoginLevelModel) error {
	args := m.Called(data)
	return  args.Error(0)
}
