package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"StorageWorkerService/pkg/logger"
)

type clientManager struct {
	Parser *JsonParser.IJsonParser
	ClientDal *abstract.IClientDal
	Log *logger.ILog
}

func ClientManagerConstructor() *clientManager {
	return &clientManager{Parser: &IoC.JsonParser,
		ClientDal: &IoC.ClientDal,
		Log: &IoC.Logger}
}

func (c *clientManager)AddClient(data *[]byte)(success bool,message string){

	client := model.ClientDataModel{}
	parseErr := (*c.Parser).DecodeJson(data, &client)
	if parseErr != nil {
		(*c.Log).SendErrorLog("ClientManager", "AddClient",
			"DecodeJson Error", parseErr.Error())
		return false, parseErr.Error()
	}

	defer (*c.Log).SendInfoLog("ClientManager", "AddClient",
		client.ClientId, client.ProjectId)

	err:= (*c.ClientDal).Add(&client)
	if err != nil {
		(*c.Log).SendErrorLog("ClientManager", "AddClient_Add",
			client.ClientId, client.ProjectId, err.Error())
		return  false, err.Error()
	}
	return  true, ""
}

func (c *clientManager)UpdateByClientId(clientId string, data *model.ClientDataModel)(success bool,message string){

	defer (*c.Log).SendInfoLog("ClientManager", "UpdateByClientId",
		clientId, data.ProjectId)

	err:= (*c.ClientDal).UpdateById(clientId, data)
	if err != nil {
		(*c.Log).SendErrorLog("ClientManager", "UpdateByClientId_UpdateById",
			clientId, data.ProjectId, err.Error())
		return  false, err.Error()
	}
	return  true, ""

}

func (c *clientManager) GetByClientId (clientId string)(data *model.ClientDataModel, success bool,message string){

	defer (*c.Log).SendInfoLog("ClientManager", "GetByClientId", clientId)

	var client, err = (*c.ClientDal).GetById(clientId)
	if err != nil {
		(*c.Log).SendErrorLog("ClientManager", "GetByClientId_GetById",
			clientId, client.ProjectId, err.Error())
	return nil, false, err.Error()
	}
	return client, true, ""
}