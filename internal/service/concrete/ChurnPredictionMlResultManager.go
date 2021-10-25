package concrete

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	jsonparser "StorageWorkerService/pkg/jsonParser"
)

type ChurnPredictionMlResultManager struct {
	Parser jsonparser.IJsonParser
	ChurnPredictionMlResultDal abstract.IChurnPredictionMlResultDal
}

func (c *ChurnPredictionMlResultManager) AddChurnPredictionMlResultData(data *[]byte)(success bool,message string){

	m := model.ChurnPredictionMlResultModel{}
	c.Parser.DecodeJson(data, &m)

	err:= c.ChurnPredictionMlResultDal.Add(&m)
	if err != nil {
		return  false, err.Error()
	}
	return  true, ""
}