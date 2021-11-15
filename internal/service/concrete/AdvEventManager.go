package concrete

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	jsonparser "StorageWorkerService/pkg/jsonParser"
)

type advEventManager struct {
	Parser jsonparser.IJsonParser
	AdvEventDal abstract.IAdvEventDal
}

func AdvEventManagerConstructor(parser jsonparser.IJsonParser, advEventDal abstract.IAdvEventDal) *advEventManager {
	return &advEventManager{Parser: parser, AdvEventDal: advEventDal}
}

func (adv *advEventManager)AddAdvEventData(data *[]byte)(success bool,message string){

		advEventDataModel := model.AdvEventDataModel{}
	adv.Parser.DecodeJson(data, &advEventDataModel)

	err:= adv.AdvEventDal.Add(&advEventDataModel)
		if err != nil {
		return  false, err.Error()
		}
		return  true, ""
}