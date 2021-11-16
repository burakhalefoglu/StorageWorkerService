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
	parseErr := (*e.Parser).DecodeJson(data, &m)
	if parseErr != nil {
		(*e.Log).SendErrorLog("EnemyBaseLevelFailManager", "AddEnemyBaseLevelFailData",
			"DecodeJson Error", parseErr.Error())
		return false, parseErr.Error()
	}
	defer (*e.Log).SendInfoLog("EnemyBaseLevelFailManager", "AddEnemyBaseLevelFailData",
		m.ClientId, m.ProjectId)

	err:= (*e.EnemyBaseLevelFailDal).Add(&m)
	if err != nil {
		(*e.Log).SendErrorLog("EnemyBaseLevelFailManager", "AddEnemyBaseLevelFailData_Add",
			m.ClientId, m.ProjectId, err.Error())

		return  false, err.Error()
	}
	return  true, ""
}