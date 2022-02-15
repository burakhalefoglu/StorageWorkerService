package abstract

import (
	"StorageWorkerService/internal/model"
)

type IOfferBehaviorDal interface {
	Add(model *model.OfferBehaviorModel) error
}
