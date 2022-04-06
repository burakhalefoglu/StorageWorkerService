package service

import (
	"StorageWorkerService/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockClientService struct {
	mock.Mock
}

func (m *MockClientService) AddClient(data *[]byte) (success bool, message string) {
	args := m.Called(data)
	return args.Bool(0), args.String(1)
}

func (m *MockClientService) UpdateByClientId(clientId int64,
	data *model.ClientDataModel) (success bool, message string) {
	args := m.Called(clientId, data)
	return args.Bool(0), args.String(1)
}

func (m *MockClientService) GetByClientId(clientId int64) (data *model.ClientDataModel, success bool, message string) {
	args := m.Called(clientId)
	return args.Get(0).(*model.ClientDataModel), args.Bool(1), args.String(2)
}
