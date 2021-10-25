package abstract

import "StorageWorkerService/internal/model"

type IScreenSwipeDal interface{
	Add(data *model.ScreenSwipeModel) error
}
