package repository

import (
	"StorageWorkerService/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockGameSessionDal struct {
	mock.Mock
}

func (m *MockGameSessionDal)Add(data *model.GameSessionModel) error {
	args := m.Called(data)
	return  args.Error(0)
}

