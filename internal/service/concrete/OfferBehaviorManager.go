package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"StorageWorkerService/pkg/logger"
)

type offerBehaviorManager struct {
	Parser *JsonParser.IJsonParser
	OfferBehaviorDal *abstract.IOfferBehaviorDal
	Log *logger.ILog
}

func OfferBehaviorManagerConstructor() *offerBehaviorManager {
	return &offerBehaviorManager{Parser: &IoC.JsonParser,
		OfferBehaviorDal: &IoC.OfferBehaviorDal,
		Log: &IoC.Logger}
}

func (o *offerBehaviorManager)AddOfferBehaviorData(data *[]byte)(success bool,message string){
	m := model.OfferBehaviorModel{}
	if err := (*o.Parser).DecodeJson(data, &m); err != nil {
		(*o.Log).SendErrorLog("OfferBehaviorManager", "AddOfferBehaviorData",
			"byte array to OfferBehaviorModel", "Json Parser Decode Err: ", err.Error())
		return false, err.Error()
	}

	defer (*o.Log).SendInfoLog("OfferBehaviorManager", "AddOfferBehaviorData",
		m.ClientId, m.ProjectId)

	if err:= (*o.OfferBehaviorDal).Add(&m); err != nil {
		(*o.Log).SendErrorLog("OfferBehaviorManager", "AddOfferBehaviorData_Add",
			"OfferBehaviorDal_Add", err.Error())
		return  false, err.Error()
	}
	return  true, ""
}

