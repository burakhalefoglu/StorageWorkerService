package concrete

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	jsonparser "StorageWorkerService/pkg/jsonParser"
)

type ScreenClickManager struct {
	Parser jsonparser.IJsonParser
	ScreenClickDal abstract.IScreenClickDal
}

func (scr *ScreenClickManager)AddScreenClickData(data *[]byte)(success bool,message string){

	screenClickData := model.ScreenClickModel{}
	scr.Parser.DecodeJson(data, &screenClickData)

	err:= scr.ScreenClickDal.Add(&screenClickData)
	if err != nil {
		return  false, err.Error()
	}
	return  true, ""
}




