package abstract

import (
	"StorageWorkerService/internal/model"
)

type IChurnPredictionSuccessRateDal interface {
	Add(model *model.ChurnPredictionSuccessRateModel) error
}
