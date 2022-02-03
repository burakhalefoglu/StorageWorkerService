package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"log"
)

type screenSwipeManager struct {
	Parser *JsonParser.IJsonParser
	ScreenSwipeDal *abstract.IScreenSwipeDal
}

func ScreenSwipeManagerConstructor() *screenSwipeManager {
	return &screenSwipeManager{Parser: &IoC.JsonParser,
		ScreenSwipeDal: &IoC.ScreenSwipeDal}
}

func (scr *screenSwipeManager)AddScreenSwipeData(data *[]byte)(success bool,message string){

	screenSwipeData := model.ScreenSwipeModel{}
	if err := (*scr.Parser).DecodeJson(data, &screenSwipeData); err != nil {
		log.Fatal("ScreenSwipeManager", "AddScreenSwipeData",
			"byte array to ScreenSwipeModel", "Json Parser Decode Err: ", err.Error())
		return false, err.Error()
	}

	defer log.Print("ScreenSwipeManager", "AddScreenSwipeData",
		screenSwipeData.ClientId, screenSwipeData.ProjectId)

	if err:= (*scr.ScreenSwipeDal).Add(&screenSwipeData); err != nil {
		log.Fatal("ScreenSwipeManager", "AddScreenSwipeData",
			"ScreenSwipeDal_Add", err.Error())
		return  false, err.Error()
	}
	return  true, ""
}



