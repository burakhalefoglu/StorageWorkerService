package concrete

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	jsonparser "StorageWorkerService/pkg/jsonParser"
)



type HardwareManager struct {
	Parser jsonparser.IJsonParser
	HardwareDal abstract.IHardwareDal
}

func (hrd *HardwareManager)AddHardwareData(data *[]byte)(success bool,message string){

	hardwareModel := model.HardwareModel{}
	hrd.Parser.DecodeJson(data, &hardwareModel)

	err:= hrd.HardwareDal.Add(&hardwareModel)
	if err != nil {
		return  false, err.Error()
	}
	return  true, ""
}