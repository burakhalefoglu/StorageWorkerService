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
	parseErr := (*i.Parser).DecodeJson(data, &m)
	if parseErr != nil {
		(*i.Log).SendErrorLog("InventoryManager", "AddInventoryData",
			"DecodeJson Error", parseErr.Error())
		return false, parseErr.Error()
	}

	defer (*i.Log).SendInfoLog("InventoryManager", "AddInventoryData",
		m.ClientId, m.ProjectId)

	err:=(*i.InventoryDal).Add(&m)
	if err != nil {
		(*i.Log).SendErrorLog("InventoryManager", "InventoryManager_Add",
			m.ClientId, m.ProjectId, err.Error())
		return  false, err.Error()
	}
	return  true, ""
}