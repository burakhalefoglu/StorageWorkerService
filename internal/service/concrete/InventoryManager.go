package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"log"
)

type inventoryManager struct {
	Parser *JsonParser.IJsonParser
	InventoryDal *abstract.IInventoryDal
}

func InventoryManagerConstructor() *inventoryManager {
	return &inventoryManager{Parser: &IoC.JsonParser,
		InventoryDal: &IoC.InventoryDal}
}

func (i *inventoryManager)AddInventoryData(data *[]byte)(success bool,message string){

	m := model.InventoryModel{}
	if err := (*i.Parser).DecodeJson(data, &m); err != nil {
		log.Fatal("InventoryManager", "AddInventoryData",
			"byte array to InventoryModel", "Json Parser Decode Err: ", err.Error())
		return false, err.Error()
	}

	defer log.Print("InventoryManager", "AddInventoryData",
		m.ClientId, m.ProjectId)

	if err:=(*i.InventoryDal).Add(&m); err != nil {
		log.Fatal("InventoryManager", "AddInventoryData",
			"InventoryDal_Add", err.Error())
		return  false, err.Error()
	}
	return  true, ""
}