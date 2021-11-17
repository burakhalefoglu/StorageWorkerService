package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"StorageWorkerService/pkg/logger"
)

type enemyBaseLoginLevelManager struct {
	Parser *JsonParser.IJsonParser
	EnemyBaseLoginLevelDal *abstract.IEnemyBaseLoginLevelDal
	Log *logger.ILog
}

func EnemyBaseLoginLevelManagerConstructor() *enemyBaseLoginLevelManager {
	return &enemyBaseLoginLevelManager{Parser: &IoC.JsonParser,
		EnemyBaseLoginLevelDal: &IoC.EnemyBaseLoginLevelDal,
		Log: &IoC.Logger}
}

func (e *enemyBaseLoginLevelManager)AddEnemyBaseLoginLevelData(data *[]byte)(success bool,message string){

	m := model.EnemyBaseLoginLevelModel{}
	if err := (*e.Parser).DecodeJson(data, &m); err != nil {
		(*e.Log).SendErrorLog("EnemyBaseLoginLevelManager", "AddEnemyBaseLoginLevelData",
			"byte array to EnemyBaseLoginLevelModel", "Json Parser Decode Err: ", err.Error())
		return false,  err.Error()
	}

	defer (*e.Log).SendInfoLog("EnemyBaseLoginLevelManager", "AddEnemyBaseLoginLevelData",
		m.ClientId, m.ProjectId)

	err:= (*e.EnemyBaseLoginLevelDal).Add(&m)
	if err != nil {
		(*e.Log).SendErrorLog("EnemyBaseLoginLevelManager", "AddEnemyBaseLoginLevelData",
			"EnemyBaseLoginLevelDal_Add", err.Error())
		return  false, err.Error()
	}
	return  true, ""
}