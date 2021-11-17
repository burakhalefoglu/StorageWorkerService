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
	if err := (*f.Parser).DecodeJson(data, &m); err != nil {
		(*f.Log).SendErrorLog("ManuelFlowManager", "AddManuelFlowData",
			"byte array to ManuelFlowModel", "Json Parser Decode Err: ", err.Error())
		return false, err.Error()
	}
	defer (*f.Log).SendInfoLog("ManuelFlowManager", "AddManuelFlowData",
		m.ClientId, m.ProjectId)

	if err:= (*f.ManuelFlowDal).Add(&m); err != nil {
		(*f.Log).SendErrorLog("ManuelFlowManager", "AddManuelFlowData",
			"ManuelFlowDal_Add", err.Error())
		return  false, err.Error()
	}
	return  true, ""
}