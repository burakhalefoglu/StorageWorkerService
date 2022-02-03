package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"log"
)

type screenClickManager struct {
	Parser *JsonParser.IJsonParser
	ScreenClickDal *abstract.IScreenClickDal
}

func ScreenClickManagerConstructor() *screenClickManager {
	return &screenClickManager{Parser: &IoC.JsonParser,
		ScreenClickDal: &IoC.ScreenClickDal}
}

func (scr *screenClickManager)AddScreenClickData(data *[]byte)(success bool,message string){

	screenClickData := model.ScreenClickModel{}
	if err := (*scr.Parser).DecodeJson(data, &screenClickData); err != nil {
		log.Fatal("ScreenClickManager", "AddScreenClickData",
			"byte array to ScreenClickModel", "Json Parser Decode Err: ", err.Error())
		return false, err.Error()
	}

	defer log.Println("ScreenClickManager", "AddScreenClickData",
		screenClickData.ClientId, screenClickData.ProjectId)

	if err:= (*scr.ScreenClickDal).Add(&screenClickData); err != nil {
		log.Fatal("ScreenClickManager", "AddScreenClickData_Add",
			"ScreenClickDal_Add", err.Error())
		return  false, err.Error()
	}
	return  true, ""
}




