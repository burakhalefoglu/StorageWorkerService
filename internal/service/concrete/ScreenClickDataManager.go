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
	if err := (*scr.Parser).DecodeJson(data, &screenClickData); err != nil {
		(*scr.Log).SendErrorLog("ScreenClickManager", "AddScreenClickData",
			"byte array to ScreenClickModel", "Json Parser Decode Err: ", err.Error())
		return false, err.Error()
	}

	defer (*scr.Log).SendInfoLog("ScreenClickManager", "AddScreenClickData",
		screenClickData.ClientId, screenClickData.ProjectId)

	if err:= (*scr.ScreenClickDal).Add(&screenClickData); err != nil {
		(*scr.Log).SendErrorLog("ScreenClickManager", "AddScreenClickData_Add",
			"ScreenClickDal_Add", err.Error())
		return  false, err.Error()
	}
	return  true, ""
}




