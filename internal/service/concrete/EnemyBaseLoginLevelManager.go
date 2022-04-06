package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
)

type enemyBaseLoginLevelManager struct {
	Parser                 *JsonParser.IJsonParser
	EnemyBaseLoginLevelDal *abstract.IEnemyBaseLoginLevelDal
}

func EnemyBaseLoginLevelManagerConstructor() *enemyBaseLoginLevelManager {
	return &enemyBaseLoginLevelManager{Parser: &IoC.JsonParser,
		EnemyBaseLoginLevelDal: &IoC.EnemyBaseLoginLevelDal}
}

func (e *enemyBaseLoginLevelManager) AddEnemyBaseLoginLevelData(data *[]byte) (success bool, message string) {

	m := model.EnemyBaseLoginLevelModel{}
	if err := (*e.Parser).DecodeJson(data, &m); err != nil {
		clogger.Error(&map[string]interface{}{
			"Json Parser Decode Err: ": err,
		})
		return false, err.Error()
	}

	defer clogger.Info(&map[string]interface{}{
		fmt.Sprintf("Data: %d", m.ClientId): "added",
	})

	err := (*e.EnemyBaseLoginLevelDal).Add(&m)
	if err != nil {
		clogger.Error(&map[string]interface{}{
			"Repository_Add Error": err,
		})
		return false, err.Error()
	}
	return true, ""
}
