package abstract

import "StorageWorkerService/internal/model"

type IAdvEventDal interface{
	Add(data *model.AdvEventDataModel) error
}
