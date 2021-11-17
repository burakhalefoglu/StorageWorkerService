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
	if err := (*c.Parser).DecodeJson(data, &churnBlockerModel); err != nil {
		(*c.Log).SendErrorLog("churnBlockerMlResultManager", "AddChurnBlockerMlResultData",
			"byte array to ChurnBlockerMlResultModel", "Json Parser Decode Err: ", err.Error())
		return false, err.Error()
	}

	defer (*c.Log).SendInfoLog("churnBlockerMlResultManager", "AddChurnBlockerMlResultData",
		churnBlockerModel.ClientId, churnBlockerModel.ProjectId)

	if err := (*c.ChurnBlockerMlResultDal).Add(&churnBlockerModel); err != nil {
		(*c.Log).SendErrorLog("churnBlockerMlResultManager", "AddChurnBlockerMlResultData",
			"ChurnBlockerMlResultDal_Add", err.Error())
		return  false, err.Error()
	}
	return  true, ""
}