package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
)

type hardwareManager struct {
	Parser      *JsonParser.IJsonParser
	HardwareDal *abstract.IHardwareDal
}

func HardwareManagerConstructor() *hardwareManager {
	return &hardwareManager{Parser: &IoC.JsonParser,
		HardwareDal: &IoC.HardwareDal}
}

func (hrd *hardwareManager) AddHardwareData(data *[]byte) (success bool, message string) {

	hardwareModel := model.HardwareModel{}
	if err := (*hrd.Parser).DecodeJson(data, &hardwareModel); err != nil {
		clogger.Error(&map[string]interface{}{
			"Json Parser Decode Err: ": err,
		})
		return false, err.Error()
	}

	defer clogger.Info(&map[string]interface{}{
		fmt.Sprintf("Data: %d", hardwareModel.ClientId): "added",
	})

	if err := (*hrd.HardwareDal).Add(&hardwareModel); err != nil {
		clogger.Error(&map[string]interface{}{
			"Repository_Add Error": err,
		})
		return false, err.Error()
	}
	return true, ""
}
