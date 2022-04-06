package abstract

import (
	"StorageWorkerService/internal/model"
)

type IAdvStrategyBehaviorDal interface {
	Add(model *model.AdvStrategyBehaviorModel) error
}
