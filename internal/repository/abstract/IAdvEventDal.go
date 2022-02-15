package abstract

import "StorageWorkerService/internal/model"

type IAdvEventDal interface {
	Add(model *model.AdvEventDataModel) error
}
