package abstract

import "StorageWorkerService/internal/model"

type IChurnPredictionMlResultDal interface{
	Add(data *model.ChurnPredictionMlResultModel) error
}
