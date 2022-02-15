package abstract

import (
	"StorageWorkerService/internal/model"
)

type IScreenSwipeDal interface {
	Add(model *model.ScreenSwipeModel) error
}
