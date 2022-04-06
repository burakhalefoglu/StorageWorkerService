package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	abstractRepo "StorageWorkerService/internal/repository/abstract"
	abstractService "StorageWorkerService/internal/service/abstract"
	JSonParser "StorageWorkerService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
)

type buyingEventManager struct {
	Parser         *JSonParser.IJsonParser
	BuyingEventDal *abstractRepo.IBuyingEventDal
	ClientService  *abstractService.IClientService
}

func BuyingEventManagerConstructor() *buyingEventManager {
	return &buyingEventManager{Parser: &IoC.JsonParser,
		BuyingEventDal: &IoC.BuyingEventDal, ClientService: &IoC.ClientService}
}

// AddBuyingEventData ! Transaction is required.
func (buying *buyingEventManager) AddBuyingEventData(data *[]byte) (success bool, message string) {

	buyingEventModel := model.BuyingEventModel{}
	if err := (*buying.Parser).DecodeJson(data, &buyingEventModel); err != nil {
		clogger.Error(&map[string]interface{}{
			"Json Parser Decode Err: ": err,
		})
		return false, err.Error()
	}

	if err := (*buying.BuyingEventDal).Add(&buyingEventModel); err != nil {
		clogger.Error(&map[string]interface{}{
			"BuyingEventDal_Add": err,
		})
		return false, err.Error()
	}

	var clientModel, s, m = (*buying.ClientService).GetByClientId(buyingEventModel.ClientId)
	if s == false {
		return false, m
	}
	if clientModel.IsPaidClient == false {
		clientModel.IsPaidClient = true
		success, message := (*buying.ClientService).UpdateByClientId(buyingEventModel.ClientId, clientModel)
		if success != true {
			return false, message
		}
	}

	defer clogger.Info(&map[string]interface{}{
		fmt.Sprintf("BuyingEvent: %d", buyingEventModel.Id): "added",
	})

	return true, ""
}
