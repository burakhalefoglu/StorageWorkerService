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

func Test_Location_SuccessIsTrue(t *testing.T) {

	//Arrange
	var testLocationDal = new(repository.MockLocationDal)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.LocationDal = testLocationDal

	location := concrete.LocationManagerConstructor()

	m:= model.LocationModel{}
	testLocationDal.On("Add", &m).Return(nil)
	message, _ := (*location.Parser).EncodeJson(&m)


	//Act
	success, err:= location.AddLocationData(message)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
}

func Test_Location_SuccessIsFalse(t *testing.T) {

	//Arrange
	var testLocationDal = new(repository.MockLocationDal)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.LocationDal = testLocationDal

	location := concrete.LocationManagerConstructor()

	m:= model.LocationModel{}
	testLocationDal.On("Add", &m).Return(errors.New("FakeError"))
	message, _ := (*location.Parser).EncodeJson(&m)


	//Act
	success, err:= location.AddLocationData(message)


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
}
