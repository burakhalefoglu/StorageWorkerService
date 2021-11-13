package abstract

import "StorageWorkerService/internal/model"

type IClientService interface{
	AddClient(data *[]byte)(success bool,message string)
	UpdateByClientId(clientId string, data *model.ClientDataModel)(success bool,message string)
	GetByClientId(clientId string)(data *model.ClientDataModel, success bool,message string)
}
