package repository

import (
	"StorageWorkerService/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockBuyingEventDal struct {
	mock.Mock
}

func (m *MockBuyingEventDal)Add(data *model.BuyingEventModel) error {
	args := m.Called(data)
	return  args.Error(0)
}
