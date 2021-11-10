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

func Test_EnemyBaseLoginLevel_SuccessIsTrue(t *testing.T) {

	//Arrange
	testObj := new(repository.MockEnemyBaseLoginLevelDal)
	enemyBaseLoginLevel := concrete.EnemyBaseLoginLevelManager{
		Parser:      &gojson.GoJson{},
		EnemyBaseLoginLevelDal: testObj,
	}
	m:= model.EnemyBaseLoginLevelModel{}
	testObj.On("Add", &m).Return(nil)
	message, _ := enemyBaseLoginLevel.Parser.EncodeJson(&m)


	//Act
	success, err:= enemyBaseLoginLevel.AddEnemyBaseLoginLevelData(message)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
}

func Test_EnemyBaseLoginLevel_SuccessIsFalse(t *testing.T) {

	//Arrange
	testObj := new(repository.MockEnemyBaseLoginLevelDal)
	enemyBaseLoginLevel:= concrete.EnemyBaseLoginLevelManager{
		Parser:      &gojson.GoJson{},
		EnemyBaseLoginLevelDal: testObj,
	}
	m:= model.EnemyBaseLoginLevelModel{}
	testObj.On("Add", &m).Return(errors.New("FakeError"))
	message, _ := enemyBaseLoginLevel.Parser.EncodeJson(&m)


	//Act
	success, err:= enemyBaseLoginLevel.AddEnemyBaseLoginLevelData(message)


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
}


