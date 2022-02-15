package abstract

import (
	"StorageWorkerService/internal/model"
)

type IEnemyBaseLevelFailDal interface {
	Add(model *model.EnemyBaseLevelFailModel) error
}
