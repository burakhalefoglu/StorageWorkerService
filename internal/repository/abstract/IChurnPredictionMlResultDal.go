package abstract

import (
	"StorageWorkerService/internal/model"
)

type IChurnPredictionMlResultDal interface {
	Add(model *model.ChurnPredictionMlResultModel) error
}
