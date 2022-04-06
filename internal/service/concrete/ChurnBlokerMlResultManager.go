package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
)

type churnBlockerMlResultManager struct {
	Parser                  *JsonParser.IJsonParser
	ChurnBlockerMlResultDal *abstract.IChurnBlockerMlResultDal
}

func ChurnBlockerMlResultManagerConstructor() *churnBlockerMlResultManager {
	return &churnBlockerMlResultManager{Parser: &IoC.JsonParser,
		ChurnBlockerMlResultDal: &IoC.ChurnBlockerMlResultDal}
}

func (c *churnBlockerMlResultManager) AddChurnBlockerMlResultData(data *[]byte) (success bool, message string) {

	churnBlockerModel := model.ChurnBlockerMlResultModel{}
	if err := (*c.Parser).DecodeJson(data, &churnBlockerModel); err != nil {
		clogger.Error(&map[string]interface{}{
			"Json Parser Decode Err: ": err,
		})
		return false, err.Error()
	}

	if err := (*c.ChurnBlockerMlResultDal).Add(&churnBlockerModel); err != nil {
		clogger.Error(&map[string]interface{}{
			"ChurnBlockerMlResultDal_Add": err,
		})
		return false, err.Error()
	}
	defer clogger.Info(&map[string]interface{}{
		fmt.Sprintf("churnBlockerMlResult: %d", churnBlockerModel.Id): "added",
	})
	return true, ""
}
