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

func Test_ScreenClick_SuccessIsTrue(t *testing.T) {

	//Arrange
	testObj := new(repository.MockScreenClickDal)
	screenClick := concrete.ScreenClickManagerConstructor(gojson.GoJsonConstructor(), testObj)

	m:= model.ScreenClickModel{}
	testObj.On("Add", &m).Return(nil)
	message, _ := screenClick.Parser.EncodeJson(&m)


	//Act
	success, err:= screenClick.AddScreenClickData(message)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
}

func Test_ScreenClick_SuccessIsFalse(t *testing.T) {

	//Arrange
	testObj := new(repository.MockScreenClickDal)
	screenClick := concrete.ScreenClickManagerConstructor(gojson.GoJsonConstructor(), testObj)

	m := model.ScreenClickModel{}
	testObj.On("Add", &m).Return(errors.New("FakeError"))
	message, _ := screenClick.Parser.EncodeJson(&m)

	//Act
	success, err := screenClick.AddScreenClickData(message)
	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
}
