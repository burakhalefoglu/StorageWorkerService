package abstract

import (
	"StorageWorkerService/internal/model"
)

type IChurnBlockerMlResultDal interface {
	Add(model *model.ChurnBlockerMlResultModel) error
}
