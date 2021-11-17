package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"StorageWorkerService/pkg/logger"
)



type hardwareManager struct {
	Parser *JsonParser.IJsonParser
	HardwareDal *abstract.IHardwareDal
	Log *logger.ILog
}

func HardwareManagerConstructor() *hardwareManager {
	return &hardwareManager{Parser: &IoC.JsonParser,
		HardwareDal: &IoC.HardwareDal,
		Log: &IoC.Logger}
}

func (hrd *hardwareManager)AddHardwareData(data *[]byte)(success bool,message string){

	hardwareModel := model.HardwareModel{}
	if err := (*hrd.Parser).DecodeJson(data, &hardwareModel); err != nil {
		(*hrd.Log).SendErrorLog("HardwareManager", "AddHardwareData",
			"byte array to HardwareModel", "Json Parser Decode Err: ", err.Error())
		return false, err.Error()
	}

	defer (*hrd.Log).SendInfoLog("HardwareManager", "AddHardwareData",
		hardwareModel.ClientId, hardwareModel.ProjectId)

	if err:= (*hrd.HardwareDal).Add(&hardwareModel); err != nil {
		(*hrd.Log).SendErrorLog("HardwareManager", "AddHardwareData",
			"HardwareDal_Add", err.Error())
		return  false, err.Error()
	}
	return  true, ""
}