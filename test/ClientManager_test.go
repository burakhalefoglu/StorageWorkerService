package test

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/service/concrete"
	"StorageWorkerService/pkg/jsonParser/gojson"
	"StorageWorkerService/test/Mocks/repository"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_AddClient_SuccessIsTrue(t *testing.T) {

	//Arrange
	var testClientDal = new(repository.MockClientDal)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.ClientDal = testClientDal

	client := concrete.ClientManagerConstructor()

	m := model.ClientDataModel{}
	testClientDal.On("Add", &m).Return(nil)
	message, _ := (*client.Parser).EncodeJson(&m)

	//Act
	success, err := client.AddClient(message)

	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
}

func Test_AddClient_SuccessIsFalse(t *testing.T) {

	//Arrange
	var testClientDal = new(repository.MockClientDal)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.ClientDal = testClientDal

	client := concrete.ClientManagerConstructor()

	m := model.ClientDataModel{}
	testClientDal.On("Add", &m).Return(errors.New("FakeError"))
	message, _ := (*client.Parser).EncodeJson(&m)

	//Act
	success, err := client.AddClient(message)

	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
}

func Test_UpdateByClientId_SuccessIsTrue(t *testing.T) {

	//Arrange
	var testClientDal = new(repository.MockClientDal)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.ClientDal = testClientDal

	client := concrete.ClientManagerConstructor()

	m := model.ClientDataModel{}
	testClientDal.On("UpdateById", 1, &m).Return(nil)

	//Act
	success, err := client.UpdateByClientId(1, &m)

	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
}

func Test_UpdateByClientId_SuccessIsFalse(t *testing.T) {

	//Arrange
	var testClientDal = new(repository.MockClientDal)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.ClientDal = testClientDal

	client := concrete.ClientManagerConstructor()

	m := model.ClientDataModel{}
	testClientDal.On("UpdateById", 1, &m).Return(errors.New("FakeError"))

	//Act
	success, err := client.UpdateByClientId(1, &m)

	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
}

func Test_GetByClientId_SuccessIsTrue(t *testing.T) {

	//Arrange
	var testClientDal = new(repository.MockClientDal)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.ClientDal = testClientDal

	client := concrete.ClientManagerConstructor()

	m := model.ClientDataModel{
		ProjectId:    1,
		IsPaidClient: false,
		CreatedAt:    time.Time{},
		PaidTime:     time.Time{},
	}
	testClientDal.On("GetById", 1).Return(&m, nil)

	//Act
	mdl, success, err := client.GetByClientId(1)

	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
	assert.Equal(t, 1, mdl.Id)
}

func Test_GetByClientId_SuccessIsFalse(t *testing.T) {

	//Arrange
	var testClientDal = new(repository.MockClientDal)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.ClientDal = testClientDal

	client := concrete.ClientManagerConstructor()

	m := model.ClientDataModel{
		ProjectId:    1,
		IsPaidClient: false,
		CreatedAt:    time.Time{},
		PaidTime:     time.Time{},
	}
	testClientDal.On("GetById", 1).Return(&m, errors.New("FakeError"))

	//Act
	mdl, success, err := client.GetByClientId(1)

	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
	assert.Nil(t, mdl)
}
