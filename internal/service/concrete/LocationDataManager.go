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
	parseErr := (*loc.Parser).DecodeJson(data, &locationModel)
	if parseErr != nil {
		(*loc.Log).SendErrorLog("locationManager", "AddLocationData",
			"DecodeJson Error", parseErr.Error())
		return false, parseErr.Error()
	}

	defer (*loc.Log).SendInfoLog("locationManager", "AddLocationData",
		locationModel.ClientId, locationModel.ProjectId)

	err:= (*loc.LocationDal).Add(&locationModel)
	if err != nil {
		(*loc.Log).SendErrorLog("locationManager", "AddLocationData_Add",
			locationModel.ClientId, locationModel.ProjectId, err.Error())
		return  false, err.Error()
	}
	return  true, ""
}