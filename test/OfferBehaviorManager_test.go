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

func Test_OfferBehavior_SuccessIsTrue(t *testing.T) {

	//Arrange
	var testOfferBehaviorDal = new(repository.MockOfferBehaviorDal)
	var json = gojson.GoJsonConstructor()
	var testLog = new(Log.MockLogger)

	IoC.JsonParser = json
	IoC.OfferBehaviorDal = testOfferBehaviorDal
	IoC.Logger = testLog

	offerBehavior := concrete.OfferBehaviorManagerConstructor()

	m:= model.OfferBehaviorModel{}
	testOfferBehaviorDal.On("Add", &m).Return(nil)
	message, _ :=(*offerBehavior.Parser).EncodeJson(&m)


	//Act
	success, err:= offerBehavior.AddOfferBehaviorData(message)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
}

func Test_OfferBehavior_SuccessIsFalse(t *testing.T) {

	//Arrange
	var testOfferBehaviorDal = new(repository.MockOfferBehaviorDal)
	var json = gojson.GoJsonConstructor()
	var testLog = new(Log.MockLogger)

	IoC.JsonParser = json
	IoC.OfferBehaviorDal = testOfferBehaviorDal
	IoC.Logger = testLog

	offerBehavior := concrete.OfferBehaviorManagerConstructor()

	m := model.OfferBehaviorModel{}
	testOfferBehaviorDal.On("Add", &m).Return(errors.New("FakeError"))
	message, _ := (*offerBehavior.Parser).EncodeJson(&m)

	//Act
	success, err := offerBehavior.AddOfferBehaviorData(message)
	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
}