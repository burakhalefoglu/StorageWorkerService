package abstract

import "StorageWorkerService/internal/model"

type IChurnBlockerMlResultDal interface{
	Add(data *model.ChurnBlockerMlResultModel) error
}
