package test

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/service/concrete"
	"StorageWorkerService/pkg/jsonParser/gojson"
	"StorageWorkerService/test/Mocks/repository"
	"StorageWorkerService/test/Mocks/service"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_BuyingEventAdd_NoClientError(t *testing.T) {

	//Arrange
	var testBuyingEventDal = new(repository.MockBuyingEventDal)
	var testClientService = new(service.MockClientService)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.BuyingEventDal = testBuyingEventDal
	IoC.ClientService = testClientService

	var buyingEvent = concrete.BuyingEventManagerConstructor()

	m:= model.BuyingEventModel{
		ClientId:      "FakeClientId",
		ProjectId:     "FakeProjectId",
		CustomerId:    "Test",
		LevelName:     "1",
		LevelIndex:    1,
		ProductType:   "test",
		InWhatMinutes: 12,
		TriggeredTime: time.Time{},
	}
	testBuyingEventDal.On("Add", &m).Return(nil)
	message, _ := (*buyingEvent.Parser).EncodeJson(&m)

	clientModel:= model.ClientDataModel{
		ClientId:     "",
		ProjectId:    "",
		IsPaidClient: 0,
		CreatedAt:    time.Time{},
		PaidTime:     time.Time{},
	}

	testClientService.On("GetByClientId", m.ClientId).Return(&clientModel, false, "NoClient")

	testClientService.On("UpdateByClientId", clientModel.ClientId,
		&clientModel).Return(true, "")

	//Act
	success, err:= buyingEvent.AddBuyingEventData(message)


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "NoClient", err)
}


func Test_BuyingEventAdd_UpdateClientError(t *testing.T) {

	//Arrange
	var testBuyingEventDal = new(repository.MockBuyingEventDal)
	var testClientService = new(service.MockClientService)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.BuyingEventDal = testBuyingEventDal
	IoC.ClientService = testClientService

	var buyingEvent = concrete.BuyingEventManagerConstructor()

	m:= model.BuyingEventModel{
		ClientId:      "FakeClientId",
		ProjectId:     "FakeProjectId",
		CustomerId:    "Test",
		LevelName:     "1",
		LevelIndex:    1,
		ProductType:   "test",
		InWhatMinutes: 12,
		TriggeredTime: time.Time{},
	}
	testBuyingEventDal.On("Add", &m).Return(nil)
	message, _ := (*buyingEvent.Parser).EncodeJson(&m)

	clientModel:= model.ClientDataModel{
		ClientId:     "FakeClientId",
		ProjectId:    "FakeProjectId",
		IsPaidClient: 0,
		CreatedAt:    time.Time{},
		PaidTime:     time.Time{},
	}

	testClientService.On("GetByClientId", m.ClientId).Return(&clientModel, true, "")

	testClientService.On("UpdateByClientId", clientModel.ClientId,
		&clientModel).Return(false, "UpdateError")

	//Act
	success, err:= buyingEvent.AddBuyingEventData(message)


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "UpdateError", err)
}



func Test_BuyingEventAdd_SuccessIsTrue(t *testing.T) {

	//Arrange
	var testBuyingEventDal = new(repository.MockBuyingEventDal)
	var testClientService = new(service.MockClientService)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.BuyingEventDal = testBuyingEventDal
	IoC.ClientService = testClientService

	var buyingEvent = concrete.BuyingEventManagerConstructor()

	m:= model.BuyingEventModel{
		ClientId:      "FakeClientId",
		ProjectId:     "FakeProjectId",
		CustomerId:    "Test",
		LevelName:     "1",
		LevelIndex:    1,
		ProductType:   "test",
		InWhatMinutes: 12,
		TriggeredTime: time.Time{},
	}
	testBuyingEventDal.On("Add", &m).Return(nil)

	message, _ := (*buyingEvent.Parser).EncodeJson(&m)

	clientModel:= model.ClientDataModel{
		ClientId:     "FakeClientId",
		ProjectId:    "FakeProjectId",
		IsPaidClient: 0,
		CreatedAt:    time.Time{},
		PaidTime:     time.Time{},
	}

	testClientService.On("GetByClientId", m.ClientId).Return(&clientModel, true, "")

	testClientService.On("UpdateByClientId", clientModel.ClientId,
		&clientModel).Return(true, "")

	//Act
	success, err:= buyingEvent.AddBuyingEventData(message)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
}

