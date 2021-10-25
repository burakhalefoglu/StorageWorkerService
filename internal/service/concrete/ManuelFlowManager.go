package concrete

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	jsonparser "StorageWorkerService/pkg/jsonParser"
)

type ManuelFlowManager struct {
	Parser jsonparser.IJsonParser
	ManuelFlowDal abstract.IManuelFlowDal
}


func (f *ManuelFlowManager)AddManuelFlowData(data *[]byte)(success bool,message string){
	m := model.ManuelFlowModel{}
	f.Parser.DecodeJson(data, &m)

	err:= f.ManuelFlowDal.Add(&m)
	if err != nil {
		return  false, err.Error()
	}
	return  true, ""
}