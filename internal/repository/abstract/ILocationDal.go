package abstract

import (
	"StorageWorkerService/internal/model"
)

type ILocationDal interface {
	Add(model *model.LocationModel) error
}
