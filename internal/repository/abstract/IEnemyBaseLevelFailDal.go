package abstract

import "StorageWorkerService/internal/model"

type IEnemyBaseLevelFailDal interface{
	Add(data *model.EnemyBaseLevelFailModel) error
}
