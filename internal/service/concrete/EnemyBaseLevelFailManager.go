package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"StorageWorkerService/pkg/logger"
)

type enemyBaseLevelFailManager struct {
	Parser *JsonParser.IJsonParser
	EnemyBaseLevelFailDal *abstract.IEnemyBaseLevelFailDal
	Log *logger.ILog
}


func EnemyBaseLevelFailManagerConstructor() *enemyBaseLevelFailManager {
	return &enemyBaseLevelFailManager{Parser: &IoC.JsonParser,
		EnemyBaseLevelFailDal: &IoC.EnemyBaseLevelFailDal,
		Log: &IoC.Logger}
}


func (e *enemyBaseLevelFailManager)AddEnemyBaseLevelFailData(data *[]byte)(success bool,message string){

	m := model.EnemyBaseLevelFailModel{}
	if err := (*e.Parser).DecodeJson(data, &m); err != nil {
		(*e.Log).SendErrorLog("EnemyBaseLevelFailManager", "AddEnemyBaseLevelFailData",
			"byte array to EnemyBaseLevelFailModel", "Json Parser Decode Err: ", err.Error())
		return false, err.Error()
	}

	defer (*e.Log).SendInfoLog("EnemyBaseLevelFailManager", "AddEnemyBaseLevelFailData",
		m.ClientId, m.ProjectId)

	if err:= (*e.EnemyBaseLevelFailDal).Add(&m); err != nil {
		(*e.Log).SendErrorLog("EnemyBaseLevelFailManager", "AddEnemyBaseLevelFailData",
			"EnemyBaseLevelFailDal_Add", err.Error())

		return  false, err.Error()
	}
	return  true, ""
}