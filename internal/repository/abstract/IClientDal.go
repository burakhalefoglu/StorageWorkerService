package abstract

import "StorageWorkerService/internal/model"

type IClientDal interface{
	Add(data *model.ClientDataModel) error
	UpdateByClientId(clientId string, data *model.ClientDataModel) error
	GetById(clientId string) (data *model.ClientDataModel, err error)

}

