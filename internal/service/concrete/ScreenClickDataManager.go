package concrete

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	jsonparser "StorageWorkerService/pkg/jsonParser"
)

type screenClickManager struct {
	Parser jsonparser.IJsonParser
	ScreenClickDal abstract.IScreenClickDal
}

func ScreenClickManagerConstructor(parser jsonparser.IJsonParser,
	screenClickDal abstract.IScreenClickDal) *screenClickManager {
	return &screenClickManager{Parser: parser, ScreenClickDal: screenClickDal}
}

func (scr *screenClickManager)AddScreenClickData(data *[]byte)(success bool,message string){

	screenClickData := model.ScreenClickModel{}
	scr.Parser.DecodeJson(data, &screenClickData)

	err:= scr.ScreenClickDal.Add(&screenClickData)
	if err != nil {
		return  false, err.Error()
	}
	return  true, ""
}




