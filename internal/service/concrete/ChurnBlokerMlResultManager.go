package concrete

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	jsonparser "StorageWorkerService/pkg/jsonParser"
)

type ChurnBlockerMlResultManager struct {
	Parser jsonparser.IJsonParser
	ChurnBlockerMlResultDal abstract.IChurnBlockerMlResultDal
}

func (c *ChurnBlockerMlResultManager)AddChurnBlockerMlResultData(data *[]byte)(success bool,message string){

	m := model.ChurnBlockerMlResultModel{}
	c.Parser.DecodeJson(data, &m)

	err:= c.ChurnBlockerMlResultDal.Add(&m)
	if err != nil {
		return  false, err.Error()
	}
	return  true, ""
}