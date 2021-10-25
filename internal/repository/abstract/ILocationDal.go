package abstract

import "StorageWorkerService/internal/model"

type ILocationDal interface{
	Add(data *model.LocationModel) error
}

