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

func Test_ManuelFlow_SuccessIsTrue(t *testing.T) {

	//Arrange
	testObj := new(repository.MockManuelFlowDal)
	manuelFlow := concrete.ManuelFlowManager{
		Parser:      &gojson.GoJson{},
		ManuelFlowDal: testObj,
	}
	m:= model.ManuelFlowModel{}
	testObj.On("Add", &m).Return(nil)
	message, _ := manuelFlow.Parser.EncodeJson(&m)


	//Act
	success, err:= manuelFlow.AddManuelFlowData(message)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
}

func Test_ManuelFlow_SuccessIsFalse(t *testing.T) {

	//Arrange
	testObj := new(repository.MockManuelFlowDal)
	manuelFlow := concrete.ManuelFlowManager{
		Parser:      &gojson.GoJson{},
		ManuelFlowDal: testObj,
	}
	m:= model.ManuelFlowModel{}
	testObj.On("Add", &m).Return(errors.New("FakeError"))
	message, _ := manuelFlow.Parser.EncodeJson(&m)


	//Act
	success, err:= manuelFlow.AddManuelFlowData(message)


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
}