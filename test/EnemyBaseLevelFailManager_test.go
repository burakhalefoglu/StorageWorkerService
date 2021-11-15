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

func Test_EnemyBaseLevelFail_SuccessIsTrue(t *testing.T) {

	//Arrange
	testObj := new(repository.MockEnemyBaseLevelFailDal)
	enemyBaseLevelFail := concrete.EnemyBaseLevelFailManagerConstructor(gojson.GoJsonConstructor(), testObj)

	m:= model.EnemyBaseLevelFailModel{}
	testObj.On("Add", &m).Return(nil)
	message, _ := enemyBaseLevelFail.Parser.EncodeJson(&m)


	//Act
	success, err:= enemyBaseLevelFail.AddEnemyBaseLevelFailData(message)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
}

func Test_EnemyBaseLevelFail_SuccessIsFalse(t *testing.T) {

	//Arrange
	testObj := new(repository.MockEnemyBaseLevelFailDal)
	enemyBaseLevelFail := concrete.EnemyBaseLevelFailManagerConstructor(gojson.GoJsonConstructor(), testObj)

	m:= model.EnemyBaseLevelFailModel{}
	testObj.On("Add", &m).Return(errors.New("FakeError"))
	message, _ := enemyBaseLevelFail.Parser.EncodeJson(&m)


	//Act
	success, err:= enemyBaseLevelFail.AddEnemyBaseLevelFailData(message)


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
}

