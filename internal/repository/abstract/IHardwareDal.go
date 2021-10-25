package abstract

import "StorageWorkerService/internal/model"

type IHardwareDal interface{
	Add(data *model.HardwareModel) error
}
