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

func Test_ScreenSwipe_SuccessIsTrue(t *testing.T) {

	//Arrange
	var testScreenSwipeDal = new(repository.MockScreenSwipeDal)
	var json = gojson.GoJsonConstructor()
	var testLog = new(Log.MockLogger)

	IoC.JsonParser = json
	IoC.ScreenSwipeDal = testScreenSwipeDal
	IoC.Logger = testLog

	screenSwipe := concrete.ScreenSwipeManagerConstructor()

	m:= model.ScreenSwipeModel{}
	testScreenSwipeDal.On("Add", &m).Return(nil)
	message, _ := (*screenSwipe.Parser).EncodeJson(&m)


	//Act
	success, err:= screenSwipe.AddScreenSwipeData(message)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
}


func Test_ScreenSwipe_SuccessIsFalse(t *testing.T) {

	//Arrange
	var testScreenSwipeDal = new(repository.MockScreenSwipeDal)
	var json = gojson.GoJsonConstructor()
	var testLog = new(Log.MockLogger)

	IoC.JsonParser = json
	IoC.ScreenSwipeDal = testScreenSwipeDal
	IoC.Logger = testLog

	screenSwipe := concrete.ScreenSwipeManagerConstructor()

	m := model.ScreenSwipeModel{}
	testScreenSwipeDal.On("Add", &m).Return(errors.New("FakeError"))
	message, _ := (*screenSwipe.Parser).EncodeJson(&m)

	//Act
	success, err := screenSwipe.AddScreenSwipeData(message)
	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
}

