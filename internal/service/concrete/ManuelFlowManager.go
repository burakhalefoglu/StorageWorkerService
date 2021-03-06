package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
)

type manuelFlowManager struct {
	Parser        *JsonParser.IJsonParser
	ManuelFlowDal *abstract.IManuelFlowDal
}

func ManuelFlowManagerConstructor() *manuelFlowManager {
	return &manuelFlowManager{Parser: &IoC.JsonParser,
		ManuelFlowDal: &IoC.ManuelFlowDal}
}

func (f *manuelFlowManager) AddManuelFlowData(data *[]byte) (success bool, message string) {
	m := model.ManuelFlowModel{}
	if err := (*f.Parser).DecodeJson(data, &m); err != nil {
		clogger.Error(&map[string]interface{}{
			"Json Parser Decode Err: ": err,
		})
		return false, err.Error()
	}
	defer clogger.Info(&map[string]interface{}{
		fmt.Sprintf("Data: %d", m.ClientId): "added",
	})

	if err := (*f.ManuelFlowDal).Add(&m); err != nil {
		clogger.Error(&map[string]interface{}{
			"Repository_Add Error": err,
		})
		return false, err.Error()
	}
	return true, ""
}
