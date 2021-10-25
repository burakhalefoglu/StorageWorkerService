package abstract

import "StorageWorkerService/internal/model"

type IScreenClickDal interface{
	Add(data *model.ScreenClickModel) error
}

