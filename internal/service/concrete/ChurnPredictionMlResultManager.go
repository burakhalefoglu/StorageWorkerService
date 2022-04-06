package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
)

type churnPredictionMlResultManager struct {
	Parser                     *JsonParser.IJsonParser
	ChurnPredictionMlResultDal *abstract.IChurnPredictionMlResultDal
}

func ChurnPredictionMlResultManagerConstructor() *churnPredictionMlResultManager {
	return &churnPredictionMlResultManager{Parser: &IoC.JsonParser,
		ChurnPredictionMlResultDal: &IoC.ChurnPredictionMlResultDal}
}

func (c *churnPredictionMlResultManager) AddChurnPredictionMlResultData(data *[]byte) (success bool, message string) {

	churnPredictionModel := model.ChurnPredictionMlResultModel{}
	if err := (*c.Parser).DecodeJson(data, &churnPredictionModel); err != nil {
		clogger.Error(&map[string]interface{}{
			"Json Parser Decode Err: ": err,
		})
		return false, err.Error()
	}

	if err := (*c.ChurnPredictionMlResultDal).Add(&churnPredictionModel); err != nil {
		clogger.Error(&map[string]interface{}{
			"ChurnPredictionMlResultDal_Add": err,
		})
		return false, err.Error()
	}

	defer clogger.Info(&map[string]interface{}{
		fmt.Sprintf("ChurnPredictionMlResult: %d", churnPredictionModel.Id): "added",
	})

	return true, ""
}
