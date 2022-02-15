package abstract

import (
	"StorageWorkerService/internal/model"
)

type IGameSessionDal interface {
	Add(model *model.GameSessionModel) error
}
