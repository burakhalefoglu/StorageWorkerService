package abstract

import "StorageWorkerService/internal/model"

type IGameSessionDal interface{
	Add(data *model.GameSessionModel) error
}

