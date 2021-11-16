package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"StorageWorkerService/pkg/logger"
)

type screenClickManager struct {
	Parser *JsonParser.IJsonParser
	ScreenClickDal *abstract.IScreenClickDal
	Log *logger.ILog
}

func ScreenClickManagerConstructor() *screenClickManager {
	return &screenClickManager{Parser: &IoC.JsonParser,
		ScreenClickDal: &IoC.ScreenClickDal,
		Log: &IoC.Logger}
}

func (scr *screenClickManager)AddScreenClickData(data *[]byte)(success bool,message string){

	screenClickData := model.ScreenClickModel{}
	parseErr := (*scr.Parser).DecodeJson(data, &screenClickData)
	if parseErr != nil {
		(*scr.Log).SendErrorLog("ScreenClickManager", "AddScreenClickData",
			"DecodeJson Error", parseErr.Error())
		return false, parseErr.Error()
	}

	defer (*scr.Log).SendInfoLog("ScreenClickManager", "AddScreenClickData",
		screenClickData.ClientId, screenClickData.ProjectId)

	err:= (*scr.ScreenClickDal).Add(&screenClickData)
	if err != nil {
		(*scr.Log).SendErrorLog("ScreenClickManager", "AddScreenClickData_Add",
			screenClickData.ClientId, screenClickData.ProjectId, err.Error())
		return  false, err.Error()
	}
	return  true, ""
}




