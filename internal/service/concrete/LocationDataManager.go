package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
)

type locationManager struct {
	Parser      *JsonParser.IJsonParser
	LocationDal *abstract.ILocationDal
}

func LocationManagerConstructor() *locationManager {
	return &locationManager{Parser: &IoC.JsonParser,
		LocationDal: &IoC.LocationDal}
}

func (loc *locationManager) AddLocationData(data *[]byte) (success bool, message string) {
	locationModel := model.LocationModel{}
	if err := (*loc.Parser).DecodeJson(data, &locationModel); err != nil {
		clogger.Error(&map[string]interface{}{
			"Json Parser Decode Err: ": err,
		})
		return false, err.Error()
	}

	defer clogger.Info(&map[string]interface{}{
		fmt.Sprintf("Data: %d", locationModel.ClientId): "added",
	})

	if err := (*loc.LocationDal).Add(&locationModel); err != nil {
		clogger.Error(&map[string]interface{}{
			"Repository_Add Error": err,
		})
		return false, err.Error()
	}
	return true, ""
}
