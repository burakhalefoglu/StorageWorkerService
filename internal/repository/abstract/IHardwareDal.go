package abstract

import (
	"StorageWorkerService/internal/model"
)

type IHardwareDal interface {
	Add(model *model.HardwareModel) error
}
