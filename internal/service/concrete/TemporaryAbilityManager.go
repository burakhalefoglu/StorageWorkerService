package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
)

type temporaryAbilityManager struct {
	Parser              *JsonParser.IJsonParser
	TemporaryAbilityDal *abstract.ITemporaryAbilityDal
}

func NewTemporaryAbilityManager() *temporaryAbilityManager {
	return &temporaryAbilityManager{Parser: &IoC.JsonParser,
		TemporaryAbilityDal: &IoC.TemporaryAbilityDal}
}

func (i *temporaryAbilityManager) AddTemporaryAbilityData(data *model.TemporaryAbilityModel) (success bool, message string) {

	defer clogger.Info(&map[string]interface{}{
		fmt.Sprintf("Data: %d", data.Id): "added",
	})

	if err := (*i.TemporaryAbilityDal).Add(data); err != nil {
		clogger.Error(&map[string]interface{}{
			"Repository_Add Error": err,
		})
		return false, err.Error()
	}
	return true, ""
}
