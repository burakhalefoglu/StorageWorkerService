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

func Test_EnemyBaseLoginLevel_SuccessIsTrue(t *testing.T) {

	//Arrange
	var testEnemyBaseLoginDal = new(repository.MockEnemyBaseLoginLevelDal)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.EnemyBaseLoginLevelDal = testEnemyBaseLoginDal

	enemyBaseLoginLevel := concrete.EnemyBaseLoginLevelManagerConstructor()

	m:= model.EnemyBaseLoginLevelModel{}
	testEnemyBaseLoginDal.On("Add", &m).Return(nil)
	message, _ := (*enemyBaseLoginLevel.Parser).EncodeJson(&m)


	//Act
	success, err:= enemyBaseLoginLevel.AddEnemyBaseLoginLevelData(message)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
}

func Test_EnemyBaseLoginLevel_SuccessIsFalse(t *testing.T) {

	//Arrange
	var testEnemyBaseLoginDal = new(repository.MockEnemyBaseLoginLevelDal)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.EnemyBaseLoginLevelDal = testEnemyBaseLoginDal

	enemyBaseLoginLevel := concrete.EnemyBaseLoginLevelManagerConstructor()

	m:= model.EnemyBaseLoginLevelModel{}
	testEnemyBaseLoginDal.On("Add", &m).Return(errors.New("FakeError"))
	message, _ := (*enemyBaseLoginLevel.Parser).EncodeJson(&m)


	//Act
	success, err:= enemyBaseLoginLevel.AddEnemyBaseLoginLevelData(message)


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
}


