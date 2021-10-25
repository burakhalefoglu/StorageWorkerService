package concrete

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	jsonparser "StorageWorkerService/pkg/jsonParser"
)

type ScreenSwipeManager struct {
	Parser jsonparser.IJsonParser
	ScreenSwipeDal abstract.IScreenSwipeDal
}

func (scr *ScreenSwipeManager)AddScreenSwipeData(data *[]byte)(success bool,message string){

	screenSwipeData := model.ScreenSwipeModel{}
	scr.Parser.DecodeJson(data, &screenSwipeData)

	err:= scr.ScreenSwipeDal.Add(&screenSwipeData)
	if err != nil {
		return  false, err.Error()
	}
	return  true, ""
}



