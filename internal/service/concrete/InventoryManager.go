package concrete

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	jsonparser "StorageWorkerService/pkg/jsonParser"
)

type inventoryManager struct {
	Parser jsonparser.IJsonParser
	InventoryDal abstract.IInventoryDal
}

func InventoryManagerConstructor(parser jsonparser.IJsonParser,
	inventoryDal abstract.IInventoryDal) *inventoryManager {
	return &inventoryManager{Parser: parser, InventoryDal: inventoryDal}
}

func (i *inventoryManager)AddInventoryData(data *[]byte)(success bool,message string){

	m := model.InventoryModel{}
	i.Parser.DecodeJson(data, &m)

	err:= i.InventoryDal.Add(&m)
	if err != nil {
		return  false, err.Error()
	}
	return  true, ""
}