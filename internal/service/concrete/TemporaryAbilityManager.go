package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"log"
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

	defer log.Print("temporaryAbilityManager", "AddTemporaryAbilityData",
		data.Id)

	if err := (*i.TemporaryAbilityDal).Add(data); err != nil {
		log.Fatal("temporaryAbilityManager", "AddTemporaryAbilityData",
			"TemporaryAbilityDal_Add", err.Error())
		return false, err.Error()
	}
	return true, ""
}
