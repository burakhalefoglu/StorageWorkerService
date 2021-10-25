package concrete

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	jsonparser "StorageWorkerService/pkg/jsonParser"
)

type InventoryManager struct {
	Parser jsonparser.IJsonParser
	InventoryDal abstract.IInventoryDal
}

func (i *InventoryManager)AddInventoryData(data *[]byte)(success bool,message string){

	m := model.InventoryModel{}
	i.Parser.DecodeJson(data, &m)

	err:= i.InventoryDal.Add(&m)
	if err != nil {
		return  false, err.Error()
	}
	return  true, ""
}