package concrete

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	jsonparser "StorageWorkerService/pkg/jsonParser"
)

type manuelFlowManager struct {
	Parser jsonparser.IJsonParser
	ManuelFlowDal abstract.IManuelFlowDal
}

func ManuelFlowManagerConstructor(parser jsonparser.IJsonParser, manuelFlowDal abstract.IManuelFlowDal) *manuelFlowManager {
	return &manuelFlowManager{Parser: parser, ManuelFlowDal: manuelFlowDal}
}

func (f *manuelFlowManager)AddManuelFlowData(data *[]byte)(success bool,message string){
	m := model.ManuelFlowModel{}
	f.Parser.DecodeJson(data, &m)

	err:= f.ManuelFlowDal.Add(&m)
	if err != nil {
		return  false, err.Error()
	}
	return  true, ""
}