package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
)

type churnPredictionSuccessRateManager struct {
	Parser                        *JsonParser.IJsonParser
	ChurnPredictionSuccessRateDal *abstract.IChurnPredictionSuccessRateDal
}

func NewChurnPredictionSuccessRate() *churnPredictionSuccessRateManager {
	return &churnPredictionSuccessRateManager{Parser: &IoC.JsonParser,
		ChurnPredictionSuccessRateDal: &IoC.ChurnPredictionSuccessRateDal}
}

func (c *churnPredictionSuccessRateManager) AddChurnPredictionSuccessRate(data *[]byte) (success bool, message string) {

	m := model.ChurnPredictionSuccessRateModel{}
	if err := (*c.Parser).DecodeJson(data, &m); err != nil {
		clogger.Error(&map[string]interface{}{
			"Json Parser Decode Err: ": err,
		})
		return false, err.Error()
	}

	defer clogger.Info(&map[string]interface{}{
		fmt.Sprintf("Data: %d", m.ProjectId): "added",
	})

	if err := (*c.ChurnPredictionSuccessRateDal).Add(&m); err != nil {
		clogger.Error(&map[string]interface{}{
			"Repository_Add Error": err,
		})
		return false, err.Error()
	}
	return true, ""
}
