package service

import (
	"StorageWorkerService/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockClientService struct {
	mock.Mock
}



func (m *MockClientService)AddClientData(data *[]byte)(success bool,message string){
	args := m.Called(data)
	return  args.Bool(0), args.String(1)
}

func (m *MockClientService)UpdateClientByClientId(clientId string,
	data *model.ClientDataModel)(success bool,message string) {
	args := m.Called(clientId, data)
	return  args.Bool(0), args.String(1)
}

func (m *MockClientService)GetByClientId(clientId string)(data *model.ClientDataModel, success bool,message string) {
	args := m.Called(clientId)
	return  args.Get(0).(*model.ClientDataModel), args.Bool(1), args.String(2)
}

