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

func Test_GameSession_SuccessIsTrue(t *testing.T) {

	//Arrange
	var testGameSessionDal = new(repository.MockGameSessionDal)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.GameSessionDal = testGameSessionDal

	gameSession := concrete.GameSessionManagerConstructor()

	m:= model.GameSessionModel{}
	testGameSessionDal.On("Add", &m).Return(nil)
	message, _ := (*gameSession.Parser).EncodeJson(&m)


	//Act
	success, err:= gameSession.AddGameSessionData(message)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
}

func Test_GameSession_SuccessIsFalse(t *testing.T) {

	//Arrange
	var testGameSessionDal = new(repository.MockGameSessionDal)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.GameSessionDal = testGameSessionDal

	gameSession := concrete.GameSessionManagerConstructor()

	m:= model.GameSessionModel{}
	testGameSessionDal.On("Add", &m).Return(errors.New("FakeError"))
	message, _ := (*gameSession.Parser).EncodeJson(&m)


	//Act
	success, err:= gameSession.AddGameSessionData(message)


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
}

