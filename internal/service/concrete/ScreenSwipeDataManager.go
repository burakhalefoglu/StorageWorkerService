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
	if err := (*scr.Parser).DecodeJson(data, &screenSwipeData); err != nil {
		(*scr.Log).SendErrorLog("ScreenSwipeManager", "AddScreenSwipeData",
			"byte array to ScreenSwipeModel", "Json Parser Decode Err: ", err.Error())
		return false, err.Error()
	}

	defer (*scr.Log).SendInfoLog("ScreenSwipeManager", "AddScreenSwipeData",
		screenSwipeData.ClientId, screenSwipeData.ProjectId)

	if err:= (*scr.ScreenSwipeDal).Add(&screenSwipeData); err != nil {
		(*scr.Log).SendErrorLog("ScreenSwipeManager", "AddScreenSwipeData",
			"ScreenSwipeDal_Add", err.Error())
		return  false, err.Error()
	}
	return  true, ""
}



