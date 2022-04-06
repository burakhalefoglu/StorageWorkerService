package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
)

type enemyBaseLevelFailManager struct {
	Parser                *JsonParser.IJsonParser
	EnemyBaseLevelFailDal *abstract.IEnemyBaseLevelFailDal
}

func EnemyBaseLevelFailManagerConstructor() *enemyBaseLevelFailManager {
	return &enemyBaseLevelFailManager{Parser: &IoC.JsonParser,
		EnemyBaseLevelFailDal: &IoC.EnemyBaseLevelFailDal}
}

func (e *enemyBaseLevelFailManager) AddEnemyBaseLevelFailData(data *[]byte) (success bool, message string) {

	m := model.EnemyBaseLevelFailModel{}
	if err := (*e.Parser).DecodeJson(data, &m); err != nil {
		clogger.Error(&map[string]interface{}{
			"Json Parser Decode Err: ": err,
		})
		return false, err.Error()
	}

	defer clogger.Info(&map[string]interface{}{
		fmt.Sprintf("Data: %d", m.ClientId): "added",
	})

	if err := (*e.EnemyBaseLevelFailDal).Add(&m); err != nil {
		clogger.Error(&map[string]interface{}{
			"Repository_Add Error": err,
		})
		return false, err.Error()
	}
	return true, ""
}
