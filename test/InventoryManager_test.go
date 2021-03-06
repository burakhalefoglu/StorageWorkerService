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

func Test_Inventory_SuccessIsTrue(t *testing.T) {

	//Arrange
	var testInventoryDal = new(repository.MockInventoryDal)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.InventoryDal = testInventoryDal

	inventoryManager := concrete.NewInventoryManager()

	m := model.InventoryModel{}
	testInventoryDal.On("Add", &m).Return(nil)
	message, _ := (*inventoryManager.Parser).EncodeJson(&m)

	//Act
	success, err := inventoryManager.AddInventoryData(message)

	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
}

func Test_Inventory_SuccessIsFalse(t *testing.T) {

	//Arrange
	var testInventoryDal = new(repository.MockInventoryDal)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.InventoryDal = testInventoryDal

	inventoryManager := concrete.NewInventoryManager()

	m := model.InventoryModel{}
	testInventoryDal.On("Add", &m).Return(errors.New("FakeError"))
	message, _ := (*inventoryManager.Parser).EncodeJson(&m)

	//Act
	success, err := inventoryManager.AddInventoryData(message)

	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
}
