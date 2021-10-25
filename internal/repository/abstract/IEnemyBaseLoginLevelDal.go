package abstract

import "StorageWorkerService/internal/model"

type IEnemyBaseLoginLevelDal interface{
	Add(data *model.EnemyBaseLoginLevelModel) error
}

