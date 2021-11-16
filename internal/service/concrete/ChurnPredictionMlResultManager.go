package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"StorageWorkerService/pkg/logger"
)

type churnPredictionMlResultManager struct {
	Parser *JsonParser.IJsonParser
	ChurnPredictionMlResultDal *abstract.IChurnPredictionMlResultDal
	Log *logger.ILog
}

func ChurnPredictionMlResultManagerConstructor() *churnPredictionMlResultManager {
	return &churnPredictionMlResultManager{Parser: &IoC.JsonParser,
		ChurnPredictionMlResultDal: &IoC.ChurnPredictionMlResultDal,
	Log: &IoC.Logger}
}

func (c *churnPredictionMlResultManager) AddChurnPredictionMlResultData(data *[]byte)(success bool,message string){

	churnPredictionModel := model.ChurnPredictionMlResultModel{}
	parseErr := (*c.Parser).DecodeJson(data, &churnPredictionModel)
	if parseErr != nil {
		(*c.Log).SendErrorLog("ChurnPredictionMlResultManager", "AddChurnPredictionMlResultData",
			"DecodeJson Error", parseErr.Error())
		return false, parseErr.Error()
	}

	defer (*c.Log).SendInfoLog("ChurnPredictionMlResultManager", "AddChurnPredictionMlResultData",
		churnPredictionModel.ClientId, churnPredictionModel.ProjectId)

	err:= (*c.ChurnPredictionMlResultDal).Add(&churnPredictionModel)
	if err != nil {
		(*c.Log).SendErrorLog("ChurnPredictionMlResultManager", "ChurnPredictionMlResultManager_Add",
			churnPredictionModel.ClientId, churnPredictionModel.ProjectId, err.Error())
		return  false, err.Error()
	}
	return  true, ""
}