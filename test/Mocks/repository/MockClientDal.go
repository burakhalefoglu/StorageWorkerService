package repository

import (
	"StorageWorkerService/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockClientDal struct {
	mock.Mock
}

func (m *MockClientDal)Add(data *model.ClientDataModel) error {
	args := m.Called(data)
	return  args.Error(0)
}

func (m *MockClientDal)UpdateByClientId(clientId string, data *model.ClientDataModel) error {
	args := m.Called(clientId, data)
	return  args.Error(0)
}

func (m *MockClientDal)GetById(clientId string) (data *model.ClientDataModel, err error){
	args := m.Called(clientId)
	return  args.Get(0).(*model.ClientDataModel), args.Error(1)
}

