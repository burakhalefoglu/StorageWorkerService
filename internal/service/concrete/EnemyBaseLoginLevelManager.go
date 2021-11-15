package concrete

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	jsonparser "StorageWorkerService/pkg/jsonParser"
)

type enemyBaseLoginLevelManager struct {
	Parser jsonparser.IJsonParser
	EnemyBaseLoginLevelDal abstract.IEnemyBaseLoginLevelDal
}

func EnemyBaseLoginLevelManagerConstructor(parser jsonparser.IJsonParser,
	enemyBaseLoginLevelDal abstract.IEnemyBaseLoginLevelDal) *enemyBaseLoginLevelManager {
	return &enemyBaseLoginLevelManager{Parser: parser, EnemyBaseLoginLevelDal: enemyBaseLoginLevelDal}
}

func (e *enemyBaseLoginLevelManager)AddEnemyBaseLoginLevelData(data *[]byte)(success bool,message string){

	m := model.EnemyBaseLoginLevelModel{}
	e.Parser.DecodeJson(data, &m)

	err:= e.EnemyBaseLoginLevelDal.Add(&m)
	if err != nil {
		return  false, err.Error()
	}
	return  true, ""
}