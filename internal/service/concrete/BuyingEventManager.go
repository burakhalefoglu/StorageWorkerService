package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	abstractRepo "StorageWorkerService/internal/repository/abstract"
	abstractService "StorageWorkerService/internal/service/abstract"
	JSonParser "StorageWorkerService/pkg/jsonParser"
	"StorageWorkerService/pkg/logger"
)

type buyingEventManager struct {
	Parser *JSonParser.IJsonParser
	BuyingEventDal *abstractRepo.IBuyingEventDal
	ClientService *abstractService.IClientService
	Log *logger.ILog

}

func BuyingEventManagerConstructor() *buyingEventManager {
	return &buyingEventManager{Parser: &IoC.JsonParser,
		BuyingEventDal: &IoC.BuyingEventDal, ClientService: &IoC.ClientService, Log: &IoC.Logger}
}

// AddBuyingEventData ! Transaction is required.
func (buying *buyingEventManager)AddBuyingEventData(data *[]byte)(success bool,message string){

	buyingEventModel := model.BuyingEventModel{}
	err := (*buying.Parser).DecodeJson(data, &buyingEventModel)
	if err != nil {
		(*buying.Log).SendErrorLog("BuyingEventManager", "AddBuyingEventData_JsonParse",
			"DecodeJson Error", err.Error())
		return false, err.Error()
	}
	defer (*buying.Log).SendInfoLog("BuyingEventManager", "AddBuyingEventData",
		buyingEventModel.ClientId, buyingEventModel.ProjectId)

	resultErr:= (*buying.BuyingEventDal).Add(&buyingEventModel)
	if resultErr != nil {
		(*buying.Log).SendErrorLog("BuyingEventManager", "BuyingEventManager_Add",
			buyingEventModel.ClientId, buyingEventModel.ProjectId, err.Error())
		return  false, resultErr.Error()
	}

	var clientModel, s, m = (*buying.ClientService).GetByClientId(buyingEventModel.ClientId)
	if s == false{
		return  false, m
	}
	if clientModel.IsPaidClient == 0{
		clientModel.IsPaidClient = 1
		success, message := (*buying.ClientService).UpdateByClientId(clientModel.ClientId, clientModel)
		if success != true{
			return  false, message
		}
	}
	return  true, ""
}