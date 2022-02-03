package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"log"
)


type locationManager struct {
	Parser *JsonParser.IJsonParser
	LocationDal *abstract.ILocationDal
}

func LocationManagerConstructor() *locationManager {
	return &locationManager{Parser: &IoC.JsonParser,
		LocationDal: &IoC.LocationDal}
}

func (loc *locationManager)AddLocationData(data *[]byte)(success bool,message string){
	locationModel := model.LocationModel{}
	if err := (*loc.Parser).DecodeJson(data, &locationModel); err != nil {
		log.Fatal("locationManager", "AddLocationData",
			"byte array to LocationModel", "Json Parser Decode Err: ", err.Error())
		return false, err.Error()
	}

	defer log.Print("locationManager", "AddLocationData",
		locationModel.ClientId, locationModel.ProjectId)

	if err := (*loc.LocationDal).Add(&locationModel); err != nil {
		log.Fatal("locationManager", "AddLocationData",
			"LocationDal_Add", err.Error())
		return  false, err.Error()
	}
	return  true, ""
}