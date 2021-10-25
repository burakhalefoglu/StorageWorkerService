package abstract

import "StorageWorkerService/internal/model"

type IOfferBehaviorDal interface{
	Add(data *model.OfferBehaviorModel) error
}
