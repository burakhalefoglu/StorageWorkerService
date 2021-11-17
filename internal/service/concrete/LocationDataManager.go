package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"StorageWorkerService/pkg/logger"
)


type locationManager struct {
	Parser *JsonParser.IJsonParser
	LocationDal *abstract.ILocationDal
	Log *logger.ILog
}

func LocationManagerConstructor() *locationManager {
	return &locationManager{Parser: &IoC.JsonParser,
		LocationDal: &IoC.LocationDal,
		Log: &IoC.Logger}
}

func (loc *locationManager)AddLocationData(data *[]byte)(success bool,message string){
	locationModel := model.LocationModel{}
	if err := (*loc.Parser).DecodeJson(data, &locationModel); err != nil {
		(*loc.Log).SendErrorLog("locationManager", "AddLocationData",
			"byte array to LocationModel", "Json Parser Decode Err: ", err.Error())
		return false, err.Error()
	}

	defer (*loc.Log).SendInfoLog("locationManager", "AddLocationData",
		locationModel.ClientId, locationModel.ProjectId)

	if err := (*loc.LocationDal).Add(&locationModel); err != nil {
		(*loc.Log).SendErrorLog("locationManager", "AddLocationData",
			"LocationDal_Add", err.Error())
		return  false, err.Error()
	}
	return  true, ""
}