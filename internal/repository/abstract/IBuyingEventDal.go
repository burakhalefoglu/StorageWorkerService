package abstract

import "StorageWorkerService/internal/model"

type IBuyingEventDal interface{
	Add(data *model.BuyingEventModel) error
}
