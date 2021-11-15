package concrete

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	jsonparser "StorageWorkerService/pkg/jsonParser"
)


type locationManager struct {
	Parser jsonparser.IJsonParser
	LocationDal abstract.ILocationDal
}

func LocationManagerConstructor(parser jsonparser.IJsonParser,
	locationDal abstract.ILocationDal) *locationManager {
	return &locationManager{Parser: parser, LocationDal: locationDal}
}

func (loc *locationManager)AddLocationData(data *[]byte)(success bool,message string){
	locationModel := model.LocationModel{}
	loc.Parser.DecodeJson(data, &locationModel)

	err:= loc.LocationDal.Add(&locationModel)
	if err != nil {
		return  false, err.Error()
	}
	return  true, ""
}