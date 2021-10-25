package concrete

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	jsonparser "StorageWorkerService/pkg/jsonParser"
)

type EnemyBaseLevelFailManager struct {
	Parser jsonparser.IJsonParser
	EnemyBaseLevelFailDal abstract.IEnemyBaseLevelFailDal
}

func (e *EnemyBaseLevelFailManager)AddEnemyBaseLevelFailData(data *[]byte)(success bool,message string){

	m := model.EnemyBaseLevelFailModel{}
	e.Parser.DecodeJson(data, &m)

	err:= e.EnemyBaseLevelFailDal.Add(&m)
	if err != nil {
		return  false, err.Error()
	}
	return  true, ""
}