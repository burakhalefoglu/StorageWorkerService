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

func Test_OfferBehavior_SuccessIsTrue(t *testing.T) {

	//Arrange
	testObj := new(repository.MockOfferBehaviorDal)
	offerBehavior := concrete.OfferBehaviorManagerConstructor(gojson.GoJsonConstructor(), testObj)

	m:= model.OfferBehaviorModel{}
	testObj.On("Add", &m).Return(nil)
	message, _ := offerBehavior.Parser.EncodeJson(&m)


	//Act
	success, err:= offerBehavior.AddOfferBehaviorData(message)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
}

func Test_OfferBehavior_SuccessIsFalse(t *testing.T) {

	//Arrange
	testObj := new(repository.MockOfferBehaviorDal)
	offerBehavior := concrete.OfferBehaviorManagerConstructor(gojson.GoJsonConstructor(), testObj)

	m := model.OfferBehaviorModel{}
	testObj.On("Add", &m).Return(errors.New("FakeError"))
	message, _ := offerBehavior.Parser.EncodeJson(&m)

	//Act
	success, err := offerBehavior.AddOfferBehaviorData(message)
	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
}