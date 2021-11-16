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
	parseErr := (*hrd.Parser).DecodeJson(data, &hardwareModel)
	if parseErr != nil {
		(*hrd.Log).SendErrorLog("HardwareManager", "AddHardwareData",
			"DecodeJson Error", parseErr.Error())
		return false, parseErr.Error()
	}

	defer (*hrd.Log).SendInfoLog("HardwareManager", "AddHardwareData",
		hardwareModel.ClientId, hardwareModel.ProjectId)

	err:= (*hrd.HardwareDal).Add(&hardwareModel)
	if err != nil {
		(*hrd.Log).SendErrorLog("HardwareManager", "AddHardwareData_Add",
			hardwareModel.ClientId, hardwareModel.ProjectId, err.Error())
		return  false, err.Error()
	}
	return  true, ""
}