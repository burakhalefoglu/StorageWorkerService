package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
)

type advEventManager struct {
	Parser      *JsonParser.IJsonParser
	AdvEventDal *abstract.IAdvEventDal
}

func AdvEventManagerConstructor() *advEventManager {
	return &advEventManager{Parser: &IoC.JsonParser, AdvEventDal: &IoC.AdvEventDal}
}

func (adv *advEventManager) AddAdvEventData(data *[]byte) (success bool, message string) {

	advEventDataModel := model.AdvEventDataModel{}
	if err := (*adv.Parser).DecodeJson(data, &advEventDataModel); err != nil {
		clogger.Error(&map[string]interface{}{
			"Json Parser Decode Err: ": err,
		})
		return false, err.Error()
	}

	if err := (*adv.AdvEventDal).Add(&advEventDataModel); err != nil {
		clogger.Error(&map[string]interface{}{
			"AdvEventDal_Add": err,
		})
		return false, err.Error()
	}

	defer clogger.Info(&map[string]interface{}{
		fmt.Sprintf("AdvEvent: %d", advEventDataModel.Id): "added",
	})

	return true, ""
}
