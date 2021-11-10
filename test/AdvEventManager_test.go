package test

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/service/concrete"
	"StorageWorkerService/pkg/jsonParser/gojson"
	"StorageWorkerService/test/Mocks/repository"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_AdvEvent_SuccessIsTrue(t *testing.T) {

	//Arrange
	testObj := new(repository.MockAdvEventDal)
	advEventManager:= concrete.AdvEventManager{
		Parser:      &gojson.GoJson{},
		AdvEventDal: testObj,
	}
	m:= model.AdvEventDataModel{}
	testObj.On("Add", &m).Return(nil)
	message, _ := advEventManager.Parser.EncodeJson(&m)


	//Act
	success, err:= advEventManager.AddAdvEventData(message)
	assert.Equal(t, "", err)


	//Assert
	assert.Equal(t, true, success)
}

func Test_AdvEvent_SuccessIsFalse(t *testing.T) {

	//Arrange
	testObj := new(repository.MockAdvEventDal)
	advEventManager:= concrete.AdvEventManager{
		Parser:      &gojson.GoJson{},
		AdvEventDal: testObj,
	}
	m:= model.AdvEventDataModel{}
	testObj.On("Add", &m).Return(errors.New("FakeError"))
	message, _ := advEventManager.Parser.EncodeJson(&m)


	//Act
	success, err:= advEventManager.AddAdvEventData(message)


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
}
