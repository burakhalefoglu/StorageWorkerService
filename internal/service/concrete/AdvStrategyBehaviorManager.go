package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
)

type advStrategyBehaviorManager struct {
	Parser                 *JsonParser.IJsonParser
	AdvStrategyBehaviorDal *abstract.IAdvStrategyBehaviorDal
}

func NewAdvStrategyBehaviorManager() *advStrategyBehaviorManager {
	return &advStrategyBehaviorManager{Parser: &IoC.JsonParser,
		AdvStrategyBehaviorDal: &IoC.AdvStrategyBehaviorDal}
}

func (o *advStrategyBehaviorManager) AddAdvStrategyBehaviorData(data *[]byte) (success bool, message string) {
	m := model.AdvStrategyBehaviorModel{}
	if err := (*o.Parser).DecodeJson(data, &m); err != nil {
		clogger.Error(&map[string]interface{}{
			"Json Parser Decode Err: ": err,
		})
		return false, err.Error()
	}

	defer clogger.Info(&map[string]interface{}{
		fmt.Sprintf("Data: %d", m.ClientId): "added",
	})

	if err := (*o.AdvStrategyBehaviorDal).Add(&m); err != nil {
		clogger.Error(&map[string]interface{}{
			"Repository_Add Error": err,
		})
		return false, err.Error()
	}
	return true, ""
}
