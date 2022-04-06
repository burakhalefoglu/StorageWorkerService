package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"

	"github.com/appneuroncompany/light-logger/clogger"
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

	defer clogger.Info(&map[string]interface{}{
		"Item Status: ": "added",
	})

	if err := (*i.ItemDal).Add(data); err != nil {
		clogger.Error(&map[string]interface{}{
			"Repository_Add Error": err,
		})
		return false, err.Error()
	}
	return true, ""
}
