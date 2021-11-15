package concrete

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	jsonparser "StorageWorkerService/pkg/jsonParser"
)

type enemyBaseLevelFailManager struct {
	Parser jsonparser.IJsonParser
	EnemyBaseLevelFailDal abstract.IEnemyBaseLevelFailDal
}

func EnemyBaseLevelFailManagerConstructor(parser jsonparser.IJsonParser,
	enemyBaseLevelFailDal abstract.IEnemyBaseLevelFailDal) *enemyBaseLevelFailManager {
	return &enemyBaseLevelFailManager{Parser: parser, EnemyBaseLevelFailDal: enemyBaseLevelFailDal}
}

func (e *enemyBaseLevelFailManager)AddEnemyBaseLevelFailData(data *[]byte)(success bool,message string){

	m := model.EnemyBaseLevelFailModel{}
	e.Parser.DecodeJson(data, &m)

	err:= e.EnemyBaseLevelFailDal.Add(&m)
	if err != nil {
		return  false, err.Error()
	}
	return  true, ""
}