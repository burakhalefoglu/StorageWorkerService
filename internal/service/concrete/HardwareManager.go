package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"log"
)



type hardwareManager struct {
	Parser *JsonParser.IJsonParser
	HardwareDal *abstract.IHardwareDal
}

func HardwareManagerConstructor() *hardwareManager {
	return &hardwareManager{Parser: &IoC.JsonParser,
		HardwareDal: &IoC.HardwareDal}
}

func (hrd *hardwareManager)AddHardwareData(data *[]byte)(success bool,message string){

	hardwareModel := model.HardwareModel{}
	if err := (*hrd.Parser).DecodeJson(data, &hardwareModel); err != nil {
		log.Fatal("HardwareManager", "AddHardwareData",
			"byte array to HardwareModel", "Json Parser Decode Err: ", err.Error())
		return false, err.Error()
	}

	defer log.Println("HardwareManager", "AddHardwareData",
		hardwareModel.ClientId, hardwareModel.ProjectId)

	if err:= (*hrd.HardwareDal).Add(&hardwareModel); err != nil {
		log.Fatal("HardwareManager", "AddHardwareData",
			"HardwareDal_Add", err.Error())
		return  false, err.Error()
	}
	return  true, ""
}