package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"StorageWorkerService/pkg/logger"
)

type screenSwipeManager struct {
	Parser *JsonParser.IJsonParser
	ScreenSwipeDal *abstract.IScreenSwipeDal
	Log *logger.ILog
}

func ScreenSwipeManagerConstructor() *screenSwipeManager {
	return &screenSwipeManager{Parser: &IoC.JsonParser,
		ScreenSwipeDal: &IoC.ScreenSwipeDal,
		Log: &IoC.Logger}
}

func (scr *screenSwipeManager)AddScreenSwipeData(data *[]byte)(success bool,message string){

	screenSwipeData := model.ScreenSwipeModel{}
	parseErr := (*scr.Parser).DecodeJson(data, &screenSwipeData)
	if parseErr != nil {
		(*scr.Log).SendErrorLog("ScreenSwipeManager", "AddScreenSwipeData",
			"DecodeJson Error", parseErr.Error())
		return false, parseErr.Error()
	}

	defer (*scr.Log).SendInfoLog("ScreenSwipeManager", "AddScreenSwipeData",
		screenSwipeData.ClientId, screenSwipeData.ProjectId)

	err:= (*scr.ScreenSwipeDal).Add(&screenSwipeData)
	if err != nil {
		(*scr.Log).SendErrorLog("ScreenSwipeManager", "AddScreenSwipeData_Add",
			screenSwipeData.ClientId, screenSwipeData.ProjectId, err.Error())
		return  false, err.Error()
	}
	return  true, ""
}



