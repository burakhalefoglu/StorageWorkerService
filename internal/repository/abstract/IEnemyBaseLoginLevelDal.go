package abstract

import (
	"StorageWorkerService/internal/model"
)

type IEnemyBaseLoginLevelDal interface {
	Add(model *model.EnemyBaseLoginLevelModel) error
}
