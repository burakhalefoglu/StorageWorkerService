package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"StorageWorkerService/pkg/logger"
)

type inventoryManager struct {
	Parser *JsonParser.IJsonParser
	InventoryDal *abstract.IInventoryDal
	Log *logger.ILog
}

func InventoryManagerConstructor() *inventoryManager {
	return &inventoryManager{Parser: &IoC.JsonParser,
		InventoryDal: &IoC.InventoryDal,
		Log: &IoC.Logger}
}

func (i *inventoryManager)AddInventoryData(data *[]byte)(success bool,message string){

	m := model.InventoryModel{}
	if err := (*i.Parser).DecodeJson(data, &m); err != nil {
		(*i.Log).SendErrorLog("InventoryManager", "AddInventoryData",
			"byte array to InventoryModel", "Json Parser Decode Err: ", err.Error())
		return false, err.Error()
	}

	defer (*i.Log).SendInfoLog("InventoryManager", "AddInventoryData",
		m.ClientId, m.ProjectId)

	if err:=(*i.InventoryDal).Add(&m); err != nil {
		(*i.Log).SendErrorLog("InventoryManager", "AddInventoryData",
			"InventoryDal_Add", err.Error())
		return  false, err.Error()
	}
	return  true, ""
}