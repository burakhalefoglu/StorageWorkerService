package repository

import (
	"StorageWorkerService/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockChurnPredictionDal struct {
	mock.Mock
}

func (m *MockChurnPredictionDal)Add(data *model.ChurnPredictionMlResultModel) error {
	args := m.Called(data)
	return  args.Error(0)
}