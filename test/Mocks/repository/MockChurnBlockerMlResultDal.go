package repository

import (
	"StorageWorkerService/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockChurnBlockerMlResultDal struct {
	mock.Mock
}

func (m *MockChurnBlockerMlResultDal)Add(data *model.ChurnBlockerMlResultModel) error {
	args := m.Called(data)
	return  args.Error(0)
}

