package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"log"
)

type enemyBaseLevelFailManager struct {
	Parser *JsonParser.IJsonParser
	EnemyBaseLevelFailDal *abstract.IEnemyBaseLevelFailDal
}


func EnemyBaseLevelFailManagerConstructor() *enemyBaseLevelFailManager {
	return &enemyBaseLevelFailManager{Parser: &IoC.JsonParser,
		EnemyBaseLevelFailDal: &IoC.EnemyBaseLevelFailDal}
}


func (e *enemyBaseLevelFailManager)AddEnemyBaseLevelFailData(data *[]byte)(success bool,message string){

	m := model.EnemyBaseLevelFailModel{}
	if err := (*e.Parser).DecodeJson(data, &m); err != nil {
		log.Fatal("EnemyBaseLevelFailManager", "AddEnemyBaseLevelFailData",
			"byte array to EnemyBaseLevelFailModel", "Json Parser Decode Err: ", err.Error())
		return false, err.Error()
	}

	defer log.Print("EnemyBaseLevelFailManager", "AddEnemyBaseLevelFailData",
		m.ClientId, m.ProjectId)

	if err:= (*e.EnemyBaseLevelFailDal).Add(&m); err != nil {
		log.Fatal("EnemyBaseLevelFailManager", "AddEnemyBaseLevelFailData",
			"EnemyBaseLevelFailDal_Add", err.Error())

		return  false, err.Error()
	}
	return  true, ""
}