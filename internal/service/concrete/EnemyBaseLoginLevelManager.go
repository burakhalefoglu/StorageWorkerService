package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"log"
)

type enemyBaseLoginLevelManager struct {
	Parser *JsonParser.IJsonParser
	EnemyBaseLoginLevelDal *abstract.IEnemyBaseLoginLevelDal
}

func EnemyBaseLoginLevelManagerConstructor() *enemyBaseLoginLevelManager {
	return &enemyBaseLoginLevelManager{Parser: &IoC.JsonParser,
		EnemyBaseLoginLevelDal: &IoC.EnemyBaseLoginLevelDal}
}

func (e *enemyBaseLoginLevelManager)AddEnemyBaseLoginLevelData(data *[]byte)(success bool,message string){

	m := model.EnemyBaseLoginLevelModel{}
	if err := (*e.Parser).DecodeJson(data, &m); err != nil {
		log.Fatal("EnemyBaseLoginLevelManager", "AddEnemyBaseLoginLevelData",
			"byte array to EnemyBaseLoginLevelModel", "Json Parser Decode Err: ", err.Error())
		return false,  err.Error()
	}

	defer log.Print("EnemyBaseLoginLevelManager", "AddEnemyBaseLoginLevelData",
		m.ClientId, m.ProjectId)

	err:= (*e.EnemyBaseLoginLevelDal).Add(&m)
	if err != nil {
		log.Fatal("EnemyBaseLoginLevelManager", "AddEnemyBaseLoginLevelData",
			"EnemyBaseLoginLevelDal_Add", err.Error())
		return  false, err.Error()
	}
	return  true, ""
}