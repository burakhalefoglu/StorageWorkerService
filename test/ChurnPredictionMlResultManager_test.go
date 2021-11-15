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

func Test_ChurnPrediction_SuccessIsTrue(t *testing.T) {

	//Arrange
	testObj := new(repository.MockChurnPredictionDal)
	churnPrediction := concrete.ChurnPredictionMlResultManagerConstructor(gojson.GoJsonConstructor(), testObj)

	m:= model.ChurnPredictionMlResultModel{}
	testObj.On("Add", &m).Return(nil)
	message, _ := churnPrediction.Parser.EncodeJson(&m)


	//Act
	success, err:= churnPrediction.AddChurnPredictionMlResultData(message)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
}

func Test_ChurnPrediction_SuccessIsFalse(t *testing.T) {

	//Arrange
	testObj := new(repository.MockChurnPredictionDal)
	churnPrediction := concrete.ChurnPredictionMlResultManagerConstructor(gojson.GoJsonConstructor(), testObj)

	m:= model.ChurnPredictionMlResultModel{}
	testObj.On("Add", &m).Return(errors.New("FakeError"))
	message, _ := churnPrediction.Parser.EncodeJson(&m)


	//Act
	success, err:= churnPrediction.AddChurnPredictionMlResultData(message)


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
}
