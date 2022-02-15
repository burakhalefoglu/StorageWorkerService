package abstract

import (
	"StorageWorkerService/internal/model"
)

type IManuelFlowDal interface {
	Add(model *model.ManuelFlowModel) error
}
