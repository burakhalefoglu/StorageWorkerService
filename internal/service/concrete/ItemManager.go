package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"log"
)

type itemManager struct {
	Parser  *JsonParser.IJsonParser
	ItemDal *abstract.IItemDal
}

func NewItemManager() *itemManager {
	return &itemManager{Parser: &IoC.JsonParser,
		ItemDal: &IoC.ItemDal}
}

func (i *itemManager) AddItemData(data *model.ItemModel) (success bool, message string) {

	defer log.Print("itemManager", "AddItemData",
		data.Id)

	if err := (*i.ItemDal).Add(data); err != nil {
		log.Fatal("itemManager", "AddItemData",
			"ItemDal_Add", err.Error())
		return false, err.Error()
	}
	return true, ""
}
