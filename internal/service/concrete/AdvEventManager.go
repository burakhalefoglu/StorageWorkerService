package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"log"
)

type advEventManager struct {
	Parser *JsonParser.IJsonParser
	AdvEventDal *abstract.IAdvEventDal
}

func AdvEventManagerConstructor() *advEventManager {
	return &advEventManager{Parser: &IoC.JsonParser, AdvEventDal: &IoC.AdvEventDal}
}

func (adv *advEventManager)AddAdvEventData(data *[]byte)(success bool,message string){

	advEventDataModel := model.AdvEventDataModel{}
	if err := (*adv.Parser).DecodeJson(data, &advEventDataModel); err != nil {
		log.Fatal("AdvEventManager", "AddAdvEventData",
			"byte array to AdvEventDataModel", "Json Parser Decode Err: ", err.Error())
		return false, err.Error()
	}

	defer log.Print("AdvEventManager", "AddAdvEventData",
		advEventDataModel.ClientId, advEventDataModel.ProjectId)

	if err:= (*adv.AdvEventDal).Add(&advEventDataModel); err != nil {
		log.Fatal("AdvEventManager", "AddAdvEventData",
				"AdvEventDal_Add", err.Error())
		return  false, err.Error()
		}

		return  true, ""
}