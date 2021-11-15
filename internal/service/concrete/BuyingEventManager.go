package concrete

import (
	"StorageWorkerService/internal/model"
	abstractRepo "StorageWorkerService/internal/repository/abstract"
	abstractService "StorageWorkerService/internal/service/abstract"
	jsonparser "StorageWorkerService/pkg/jsonParser"
)

type buyingEventManager struct {
	Parser jsonparser.IJsonParser
	BuyingEventDal abstractRepo.IBuyingEventDal
	ClientService abstractService.IClientService
}

func BuyingEventManagerConstructor(parser jsonparser.IJsonParser,
	buyingEventDal abstractRepo.IBuyingEventDal,
	clientService abstractService.IClientService) *buyingEventManager {
	return &buyingEventManager{Parser: parser, BuyingEventDal: buyingEventDal, ClientService: clientService}
}

//! Transaction gerekecek.
func (buying *buyingEventManager)AddBuyingEventData(data *[]byte)(success bool,message string){

	buyingEventModel := model.BuyingEventModel{}
	buying.Parser.DecodeJson(data, &buyingEventModel)

	resultErr:= buying.BuyingEventDal.Add(&buyingEventModel)
	if resultErr != nil {
		return  false, resultErr.Error()
	}
	var clientModel, s, m = buying.ClientService.GetByClientId(buyingEventModel.ClientId)
	if s == false{
		return  false, m
	}
	if clientModel.IsPaidClient == 0{
		clientModel.IsPaidClient = 1
		success, message := buying.ClientService.UpdateByClientId(clientModel.ClientId, clientModel)
		if success != true{
			return  false, message
		}
	}
	return  true, ""
}