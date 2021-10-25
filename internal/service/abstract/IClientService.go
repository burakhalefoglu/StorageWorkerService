package abstract

import "StorageWorkerService/internal/model"

type IClientService interface{
	AddClientData(data *[]byte)(success bool,message string)
	UpdateClientByClientId(clientId string, data *model.ClientDataModel)(success bool,message string)
	GetByClientId(clientId string)(data *model.ClientDataModel, success bool,message string)
}
