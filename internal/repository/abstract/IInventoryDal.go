package abstract

import "StorageWorkerService/internal/model"

type IInventoryDal interface{
	Add(data *model.InventoryModel) error
}

