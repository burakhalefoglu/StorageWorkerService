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

func Test_GameSession_SuccessIsTrue(t *testing.T) {

	//Arrange
	testObj := new(repository.MockGameSessionDal)
	gameSession := concrete.GameSessionManager{
		Parser:      &gojson.GoJson{},
		GameSessionDal: testObj,
	}
	m:= model.GameSessionModel{}
	testObj.On("Add", &m).Return(nil)
	message, _ := gameSession.Parser.EncodeJson(&m)


	//Act
	success, err:= gameSession.AddGameSessionData(message)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
}

func Test_GameSession_SuccessIsFalse(t *testing.T) {

	//Arrange
	testObj := new(repository.MockGameSessionDal)
	gameSession:= concrete.GameSessionManager{
		Parser:      &gojson.GoJson{},
		GameSessionDal: testObj,
	}
	m:= model.GameSessionModel{}
	testObj.On("Add", &m).Return(errors.New("FakeError"))
	message, _ := gameSession.Parser.EncodeJson(&m)


	//Act
	success, err:= gameSession.AddGameSessionData(message)


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
}

