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

func Test_LevelBaseSession_SuccessIsTrue(t *testing.T) {

	//Arrange
	var levelBaseSessionDal = new(repository.MockLevelBaseSessionDal)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.LevelBaseSessionDal = levelBaseSessionDal

	levelBaseSession := concrete.LevelBaseSessionManagerConstructor()

	m:= model.LevelBaseSessionModel{}
	levelBaseSessionDal.On("Add", &m).Return(nil)
	message, _ := (*levelBaseSession.Parser).EncodeJson(&m)


	//Act
	success, err:= levelBaseSession.AddLevelBaseSessionData(message)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
}

func Test_LevelBaseSession_SuccessIsFalse(t *testing.T) {

	//Arrange
	var levelBaseSessionDal = new(repository.MockLevelBaseSessionDal)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.LevelBaseSessionDal = levelBaseSessionDal

	levelBaseSession := concrete.LevelBaseSessionManagerConstructor()

	m:= model.LevelBaseSessionModel{}
	levelBaseSessionDal.On("Add", &m).Return(errors.New("FakeError"))
	message, _ := (*levelBaseSession.Parser).EncodeJson(&m)


	//Act
	success, err:= levelBaseSession.AddLevelBaseSessionData(message)


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
}

