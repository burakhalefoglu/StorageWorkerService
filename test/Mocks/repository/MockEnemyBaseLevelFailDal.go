package repository

import (
	"StorageWorkerService/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockEnemyBaseLevelFailDal struct {
	mock.Mock
}

func (m *MockEnemyBaseLevelFailDal)Add(data *model.EnemyBaseLevelFailModel) error {
	args := m.Called(data)
	return  args.Error(0)
}
