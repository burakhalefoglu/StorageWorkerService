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
)

func Test_AddAdvEvent_SuccessIsTrue(t *testing.T) {

	//Arrange
	var testAdvDal = new(repository.MockAdvEventDal)
	var json = gojson.GoJsonConstructor()
	var testLog = new(Log.MockLogger)

	IoC.JsonParser = json
	IoC.AdvEventDal = testAdvDal
	IoC.Logger = testLog

	advEventManager:= concrete.AdvEventManagerConstructor()

	m:= model.AdvEventDataModel{}
	testAdvDal.On("Add", &m).Return(nil)
	message, _ := (*advEventManager.Parser).EncodeJson(&m)


	//Act
	success, err:= advEventManager.AddAdvEventData(message)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)

}

func Test_AddAdvEvent_SuccessIsFalse(t *testing.T) {

	//Arrange
	mockAdvDal := new(repository.MockAdvEventDal)
	json := gojson.GoJsonConstructor()
	mockLog :=new(Log.MockLogger)

	IoC.JsonParser = json
	IoC.AdvEventDal = mockAdvDal
	IoC.Logger = mockLog

	advEventManager:= concrete.AdvEventManagerConstructor()

	m:= model.AdvEventDataModel{}
	mockAdvDal.On("Add", &m).Return(errors.New("FakeError"))
	message, _ := (*advEventManager.Parser).EncodeJson(&m)


	//Act
	success, err:= advEventManager.AddAdvEventData(message)


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
}
