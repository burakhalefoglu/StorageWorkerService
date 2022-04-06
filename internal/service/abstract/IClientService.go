package abstract

import "StorageWorkerService/internal/model"

type IClientService interface {
	AddClient(data *[]byte) (success bool, message string)
	UpdateByClientId(clientId int64, data *model.ClientDataModel) (success bool, message string)
	GetByClientId(clientId int64) (data *model.ClientDataModel, success bool, message string)
}
