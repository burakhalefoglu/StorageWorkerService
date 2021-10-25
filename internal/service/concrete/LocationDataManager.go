package concrete

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	jsonparser "StorageWorkerService/pkg/jsonParser"
)


type LocationManager struct {
	Parser jsonparser.IJsonParser
	LocationDal abstract.ILocationDal
}


func (loc *LocationManager)AddLocationData(data *[]byte)(success bool,message string){
	locationModel := model.LocationModel{}
	loc.Parser.DecodeJson(data, &locationModel)

	err:= loc.LocationDal.Add(&locationModel)
	if err != nil {
		return  false, err.Error()
	}
	return  true, ""
}