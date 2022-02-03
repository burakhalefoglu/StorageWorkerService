package test

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/service/concrete"
	"StorageWorkerService/pkg/jsonParser/gojson"
	"StorageWorkerService/test/Mocks/repository"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ScreenClick_SuccessIsTrue(t *testing.T) {

	//Arrange
	var testScreenClickDal = new(repository.MockScreenClickDal)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.ScreenClickDal = testScreenClickDal

	screenClick := concrete.ScreenClickManagerConstructor()

	m:= model.ScreenClickModel{}
	testScreenClickDal.On("Add", &m).Return(nil)
	message, _ :=(*screenClick.Parser).EncodeJson(&m)


	//Act
	success, err:= screenClick.AddScreenClickData(message)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
}

func Test_ScreenClick_SuccessIsFalse(t *testing.T) {

	//Arrange
	var testScreenClickDal = new(repository.MockScreenClickDal)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.ScreenClickDal = testScreenClickDal

	screenClick := concrete.ScreenClickManagerConstructor()

	m := model.ScreenClickModel{}
	testScreenClickDal.On("Add", &m).Return(errors.New("FakeError"))
	message, _ := (*screenClick.Parser).EncodeJson(&m)

	//Act
	success, err := screenClick.AddScreenClickData(message)
	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
}
