package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
)

type offerBehaviorManager struct {
	Parser           *JsonParser.IJsonParser
	OfferBehaviorDal *abstract.IOfferBehaviorDal
}

func OfferBehaviorManagerConstructor() *offerBehaviorManager {
	return &offerBehaviorManager{Parser: &IoC.JsonParser,
		OfferBehaviorDal: &IoC.OfferBehaviorDal}
}

func (o *offerBehaviorManager) AddOfferBehaviorData(data *[]byte) (success bool, message string) {
	m := model.OfferBehaviorModel{}
	if err := (*o.Parser).DecodeJson(data, &m); err != nil {
		clogger.Error(&map[string]interface{}{
			"Json Parser Decode Err: ": err,
		})
		return false, err.Error()
	}

	defer clogger.Info(&map[string]interface{}{
		fmt.Sprintf("Data: %d", m.ClientId): "added",
	})

	if err := (*o.OfferBehaviorDal).Add(&m); err != nil {
		clogger.Error(&map[string]interface{}{
			"Repository_Add Error": err,
		})
		return false, err.Error()
	}
	return true, ""
}
