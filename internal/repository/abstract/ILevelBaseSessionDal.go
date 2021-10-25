package abstract

import "StorageWorkerService/internal/model"

type ILevelBaseSessionDal interface{
	Add(data *model.LevelBaseSessionModel) error
}
