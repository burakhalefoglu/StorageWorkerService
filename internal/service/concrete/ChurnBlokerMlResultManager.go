package concrete

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	jsonparser "StorageWorkerService/pkg/jsonParser"
)

type churnBlockerMlResultManager struct {
	Parser jsonparser.IJsonParser
	ChurnBlockerMlResultDal abstract.IChurnBlockerMlResultDal
}

func ChurnBlockerMlResultManagerConstructor(parser jsonparser.IJsonParser,
	churnBlockerMlResultDal abstract.IChurnBlockerMlResultDal) *churnBlockerMlResultManager {
	return &churnBlockerMlResultManager{Parser: parser, ChurnBlockerMlResultDal: churnBlockerMlResultDal}
}

func (c *churnBlockerMlResultManager)AddChurnBlockerMlResultData(data *[]byte)(success bool,message string){

	m := model.ChurnBlockerMlResultModel{}
	c.Parser.DecodeJson(data, &m)

	err:= c.ChurnBlockerMlResultDal.Add(&m)
	if err != nil {
		return  false, err.Error()
	}
	return  true, ""
}