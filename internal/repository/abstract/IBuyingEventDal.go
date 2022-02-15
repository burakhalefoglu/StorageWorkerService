package abstract

import (
	"StorageWorkerService/internal/model"
)

type IBuyingEventDal interface {
	Add(model *model.BuyingEventModel) error
}
