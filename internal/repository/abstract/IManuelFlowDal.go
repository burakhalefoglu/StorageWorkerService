package abstract

import "StorageWorkerService/internal/model"

type IManuelFlowDal interface{
	Add(data *model.ManuelFlowModel) error
}

