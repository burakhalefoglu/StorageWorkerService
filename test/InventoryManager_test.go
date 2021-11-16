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

func Test_Inventory_SuccessIsTrue(t *testing.T) {

	//Arrange
	var testInventoryDal = new(repository.MockInventoryDal)
	var json = gojson.GoJsonConstructor()
	var testLog = new(Log.MockLogger)

	IoC.JsonParser = json
	IoC.InventoryDal = testInventoryDal
	IoC.Logger = testLog

	inventoryManager := concrete.InventoryManagerConstructor()

	m:= model.InventoryModel{}
	testInventoryDal.On("Add", &m).Return(nil)
	message, _ := (*inventoryManager.Parser).EncodeJson(&m)


	//Act
	success, err:= inventoryManager.AddInventoryData(message)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
}

func Test_Inventory_SuccessIsFalse(t *testing.T) {

	//Arrange
	var testInventoryDal = new(repository.MockInventoryDal)
	var json = gojson.GoJsonConstructor()
	var testLog = new(Log.MockLogger)

	IoC.JsonParser = json
	IoC.InventoryDal = testInventoryDal
	IoC.Logger = testLog

	inventoryManager := concrete.InventoryManagerConstructor()

	m:= model.InventoryModel{}
	testInventoryDal.On("Add", &m).Return(errors.New("FakeError"))
	message, _ := (*inventoryManager.Parser).EncodeJson(&m)


	//Act
	success, err:= inventoryManager.AddInventoryData(message)


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
}

