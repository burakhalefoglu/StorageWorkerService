package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"StorageWorkerService/pkg/logger"
)

type churnBlockerMlResultManager struct {
	Parser *JsonParser.IJsonParser
	ChurnBlockerMlResultDal *abstract.IChurnBlockerMlResultDal
	Log *logger.ILog
}

func ChurnBlockerMlResultManagerConstructor() *churnBlockerMlResultManager {
	return &churnBlockerMlResultManager{Parser: &IoC.JsonParser,
		ChurnBlockerMlResultDal: &IoC.ChurnBlockerMlResultDal,
		Log: &IoC.Logger}
}

func (c *churnBlockerMlResultManager)AddChurnBlockerMlResultData(data *[]byte)(success bool,message string){

	churnBlockerModel := model.ChurnBlockerMlResultModel{}
	parseErr := (*c.Parser).DecodeJson(data, &churnBlockerModel)
	if parseErr != nil {
		(*c.Log).SendErrorLog("churnBlockerMlResultManager", "AddChurnBlockerMlResultData",
			"DecodeJson Error", parseErr.Error())
		return false, parseErr.Error()
	}
	defer (*c.Log).SendInfoLog("churnBlockerMlResultManager", "AddChurnBlockerMlResultData",
		churnBlockerModel.ClientId, churnBlockerModel.ProjectId)

	err:= (*c.ChurnBlockerMlResultDal).Add(&churnBlockerModel)
	if err != nil {
		(*c.Log).SendErrorLog("churnBlockerMlResultManager", "AddAdvEventData_Add",
			churnBlockerModel.ClientId, churnBlockerModel.ProjectId, err.Error())
		return  false, err.Error()
	}
	return  true, ""
}