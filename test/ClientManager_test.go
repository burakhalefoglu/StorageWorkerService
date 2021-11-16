package test

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/service/concrete"
	"StorageWorkerService/pkg/jsonParser/gojson"
	"StorageWorkerService/test/Mocks/Log"
	"StorageWorkerService/test/Mocks/repository"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_AddClient_SuccessIsTrue(t *testing.T) {

	//Arrange
	var testClientDal = new(repository.MockClientDal)
	var json = gojson.GoJsonConstructor()
	var testLog = new(Log.MockLogger)

	IoC.JsonParser = json
	IoC.ClientDal = testClientDal
	IoC.Logger = testLog

	client := concrete.ClientManagerConstructor()

	m:= model.ClientDataModel{}
	testClientDal.On("Add", &m).Return(nil)
	message, _ := (*client.Parser).EncodeJson(&m)


	//Act
	success, err:= client.AddClient(message)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
}

func Test_AddClient_SuccessIsFalse(t *testing.T) {

	//Arrange
	var testClientDal = new(repository.MockClientDal)
	var json = gojson.GoJsonConstructor()
	var testLog = new(Log.MockLogger)

	IoC.JsonParser = json
	IoC.ClientDal = testClientDal
	IoC.Logger = testLog

	client := concrete.ClientManagerConstructor()

	m:= model.ClientDataModel{}
	testClientDal.On("Add", &m).Return(errors.New("FakeError"))
	message, _ := (*client.Parser).EncodeJson(&m)


	//Act
	success, err:= client.AddClient(message)


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
}

func Test_UpdateByClientId_SuccessIsTrue(t *testing.T) {

	//Arrange
	var testClientDal = new(repository.MockClientDal)
	var json = gojson.GoJsonConstructor()
	var testLog = new(Log.MockLogger)

	IoC.JsonParser = json
	IoC.ClientDal = testClientDal
	IoC.Logger = testLog

	client := concrete.ClientManagerConstructor()

	m:= model.ClientDataModel{}
	testClientDal.On("UpdateById","fakeClientId", &m).Return(nil)

	//Act
	success, err:= client.UpdateByClientId("fakeClientId", &m)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
}

func Test_UpdateByClientId_SuccessIsFalse(t *testing.T) {

	//Arrange
	var testClientDal = new(repository.MockClientDal)
	var json = gojson.GoJsonConstructor()
	var testLog = new(Log.MockLogger)

	IoC.JsonParser = json
	IoC.ClientDal = testClientDal
	IoC.Logger = testLog

	client := concrete.ClientManagerConstructor()

	m:= model.ClientDataModel{}
	testClientDal.On("UpdateById","fakeClientId", &m).Return(errors.New("FakeError"))

	//Act
	success, err:= client.UpdateByClientId("fakeClientId", &m)


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
}

func Test_GetByClientId_SuccessIsTrue(t *testing.T) {

	//Arrange
	var testClientDal = new(repository.MockClientDal)
	var json = gojson.GoJsonConstructor()
	var testLog = new(Log.MockLogger)

	IoC.JsonParser = json
	IoC.ClientDal = testClientDal
	IoC.Logger = testLog

	client := concrete.ClientManagerConstructor()

	m:= model.ClientDataModel{
		ClientId:     "FakeClientId",
		ProjectId:    "FakeProjectId",
		IsPaidClient: 0,
		CreatedAt:    time.Time{},
		PaidTime:     time.Time{},
	}
	testClientDal.On("GetById","fakeClientId").Return(&m, nil)

	//Act
	mdl,success, err:= client.GetByClientId("fakeClientId")


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
	assert.Equal(t, "FakeClientId", mdl.ClientId)
}

func Test_GetByClientId_SuccessIsFalse(t *testing.T) {

	//Arrange
	var testClientDal = new(repository.MockClientDal)
	var json = gojson.GoJsonConstructor()
	var testLog = new(Log.MockLogger)

	IoC.JsonParser = json
	IoC.ClientDal = testClientDal
	IoC.Logger = testLog

	client := concrete.ClientManagerConstructor()

	m:= model.ClientDataModel{
		ClientId:     "FakeClientId",
		ProjectId:    "FakeProjectId",
		IsPaidClient: 0,
		CreatedAt:    time.Time{},
		PaidTime:     time.Time{},
	}
	testClientDal.On("GetById","fakeClientId").Return(&m, errors.New("FakeError"))

	//Act
	mdl,success, err:= client.GetByClientId("fakeClientId")


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
	assert.Nil(t, mdl)
}