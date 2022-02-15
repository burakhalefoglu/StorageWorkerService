package abstract

import (
	"StorageWorkerService/internal/model"
)

type IScreenClickDal interface {
	Add(model *model.ScreenClickModel) error
}
