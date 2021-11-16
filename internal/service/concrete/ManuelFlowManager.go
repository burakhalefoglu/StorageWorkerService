package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"StorageWorkerService/pkg/logger"
)

type manuelFlowManager struct {
	Parser *JsonParser.IJsonParser
	ManuelFlowDal *abstract.IManuelFlowDal
	Log *logger.ILog
}

func ManuelFlowManagerConstructor() *manuelFlowManager {
	return &manuelFlowManager{Parser: &IoC.JsonParser,
		ManuelFlowDal: &IoC.ManuelFlowDal,
		Log: &IoC.Logger}
}

func (f *manuelFlowManager)AddManuelFlowData(data *[]byte)(success bool,message string){
	m := model.ManuelFlowModel{}
	parseErr := (*f.Parser).DecodeJson(data, &m)
	if parseErr != nil {
		(*f.Log).SendErrorLog("ManuelFlowManager", "AddManuelFlowData",
			"DecodeJson Error", parseErr.Error())
		return false, parseErr.Error()
	}
	defer (*f.Log).SendInfoLog("ManuelFlowManager", "AddManuelFlowData",
		m.ClientId, m.ProjectId)

	err:= (*f.ManuelFlowDal).Add(&m)
	if err != nil {
		(*f.Log).SendErrorLog("ManuelFlowManager", "AddManuelFlowData_Add",
			m.ClientId, m.ProjectId, err.Error())
		return  false, err.Error()
	}
	return  true, ""
}