package abstract

import (
	"StorageWorkerService/internal/model"
)

type ILevelBaseSessionDal interface {
	Add(model *model.LevelBaseSessionModel) error
}
