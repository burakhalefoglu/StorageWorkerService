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

func Test_LevelBaseSession_SuccessIsTrue(t *testing.T) {

	//Arrange
	testObj := new(repository.MockLevelBaseSessionDal)
	levelBaseSession := concrete.LevelBaseSessionManagerConstructor(gojson.GoJsonConstructor(), testObj)

	m:= model.LevelBaseSessionModel{}
	testObj.On("Add", &m).Return(nil)
	message, _ := levelBaseSession.Parser.EncodeJson(&m)


	//Act
	success, err:= levelBaseSession.AddLevelBaseSessionData(message)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
}

func Test_LevelBaseSession_SuccessIsFalse(t *testing.T) {

	//Arrange
	testObj := new(repository.MockLevelBaseSessionDal)
	levelBaseSession := concrete.LevelBaseSessionManagerConstructor(gojson.GoJsonConstructor(), testObj)

	m:= model.LevelBaseSessionModel{}
	testObj.On("Add", &m).Return(errors.New("FakeError"))
	message, _ := levelBaseSession.Parser.EncodeJson(&m)


	//Act
	success, err:= levelBaseSession.AddLevelBaseSessionData(message)


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
}

