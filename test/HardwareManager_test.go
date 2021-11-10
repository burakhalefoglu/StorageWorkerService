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

func Test_Hardware_SuccessIsTrue(t *testing.T) {

	//Arrange
	testObj := new(repository.MockHardwareDal)
	hardware := concrete.HardwareManager{
		Parser:      &gojson.GoJson{},
		HardwareDal: testObj,
	}
	m:= model.HardwareModel{}
	testObj.On("Add", &m).Return(nil)
	message, _ := hardware.Parser.EncodeJson(&m)


	//Act
	success, err:= hardware.AddHardwareData(message)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
}

func Test_Hardware_SuccessIsFalse(t *testing.T) {

	//Arrange
	testObj := new(repository.MockHardwareDal)
	hardware := concrete.HardwareManager{
		Parser:      &gojson.GoJson{},
		HardwareDal: testObj,
	}
	m:= model.HardwareModel{}
	testObj.On("Add", &m).Return(errors.New("FakeError"))
	message, _ := hardware.Parser.EncodeJson(&m)


	//Act
	success, err:= hardware.AddHardwareData(message)


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
}
