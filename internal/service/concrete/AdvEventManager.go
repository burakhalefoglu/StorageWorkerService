package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"StorageWorkerService/pkg/logger"
)

type advEventManager struct {
	Parser *JsonParser.IJsonParser
	AdvEventDal *abstract.IAdvEventDal
	Log *logger.ILog
}

func AdvEventManagerConstructor() *advEventManager {
	return &advEventManager{Parser: &IoC.JsonParser, AdvEventDal: &IoC.AdvEventDal, Log: &IoC.Logger}
}

func (adv *advEventManager)AddAdvEventData(data *[]byte)(success bool,message string){

	advEventDataModel := model.AdvEventDataModel{}
	parseErr := (*adv.Parser).DecodeJson(data, &advEventDataModel)
	if parseErr != nil {
		(*adv.Log).SendErrorLog("AdvEventManager", "AddAdvEventData",
			"DecodeJson Error", parseErr.Error())
		return false, parseErr.Error()
	}

	defer (*adv.Log).SendInfoLog("AdvEventManager", "AddAdvEventData",
		advEventDataModel.ClientId, advEventDataModel.ProjectId)

	err:= (*adv.AdvEventDal).Add(&advEventDataModel)
		if err != nil {
			(*adv.Log).SendErrorLog("AdvEventManager", "AddAdvEventData_Add",
				advEventDataModel.ClientId, advEventDataModel.ProjectId, err.Error())
		return  false, err.Error()
		}

		return  true, ""
}