package concrete

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	jsonparser "StorageWorkerService/pkg/jsonParser"
)



type hardwareManager struct {
	Parser jsonparser.IJsonParser
	HardwareDal abstract.IHardwareDal
}

func HardwareManagerConstructor(parser jsonparser.IJsonParser,
	hardwareDal abstract.IHardwareDal) *hardwareManager {
	return &hardwareManager{Parser: parser, HardwareDal: hardwareDal}
}

func (hrd *hardwareManager)AddHardwareData(data *[]byte)(success bool,message string){

	hardwareModel := model.HardwareModel{}
	hrd.Parser.DecodeJson(data, &hardwareModel)

	err:= hrd.HardwareDal.Add(&hardwareModel)
	if err != nil {
		return  false, err.Error()
	}
	return  true, ""
}