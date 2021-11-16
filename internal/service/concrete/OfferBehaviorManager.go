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
	parseErr := (*o.Parser).DecodeJson(data, &m)
	if parseErr != nil {
		(*o.Log).SendErrorLog("OfferBehaviorManager", "AddOfferBehaviorData",
			"DecodeJson Error", parseErr.Error())
		return false, parseErr.Error()
	}

	defer (*o.Log).SendInfoLog("OfferBehaviorManager", "AddOfferBehaviorData",
		m.ClientId, m.ProjectId)

	err:= (*o.OfferBehaviorDal).Add(&m)
	if err != nil {
		(*o.Log).SendErrorLog("OfferBehaviorManager", "AddOfferBehaviorData_Add",
			m.ClientId, m.ProjectId, err.Error())
		return  false, err.Error()
	}
	return  true, ""
}

