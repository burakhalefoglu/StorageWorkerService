package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	abstractRepo "StorageWorkerService/internal/repository/abstract"
	abstractService "StorageWorkerService/internal/service/abstract"
	JSonParser "StorageWorkerService/pkg/jsonParser"
	"log"
)

type buyingEventManager struct {
	Parser *JSonParser.IJsonParser
	BuyingEventDal *abstractRepo.IBuyingEventDal
	ClientService *abstractService.IClientService
}

func BuyingEventManagerConstructor() *buyingEventManager {
	return &buyingEventManager{Parser: &IoC.JsonParser,
		BuyingEventDal: &IoC.BuyingEventDal, ClientService: &IoC.ClientService}
}

// AddBuyingEventData ! Transaction is required.
func (buying *buyingEventManager)AddBuyingEventData(data *[]byte)(success bool,message string){

	buyingEventModel := model.BuyingEventModel{}
	if err := (*buying.Parser).DecodeJson(data, &buyingEventModel); err != nil {
		log.Fatal("BuyingEventManager", "AddAdvEventData",
			"byte array to BuyingEventModel", "Json Parser Decode Err: ", err.Error())
		return false, err.Error()
	}

	defer log.Print("BuyingEventManager", "AddBuyingEventData",
		buyingEventModel.ClientId, buyingEventModel.ProjectId)

	if err := (*buying.BuyingEventDal).Add(&buyingEventModel); err != nil {
		log.Fatal("BuyingEventManager", "AddBuyingEventData",
			"BuyingEventDal_Add", err.Error())
		return  false, err.Error()
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