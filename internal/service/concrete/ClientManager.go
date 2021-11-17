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
	if err := (*c.Parser).DecodeJson(data, &client); err != nil {
		(*c.Log).SendErrorLog("ClientManager", "AddClient",
			"byte array to ClientDataModel", "Json Parser Decode Err: ", err.Error())
		return false, err.Error()
	}

	defer (*c.Log).SendInfoLog("ClientManager", "AddClient",
		client.ClientId, client.ProjectId)

	if err:= (*c.ClientDal).Add(&client); err != nil {
		(*c.Log).SendErrorLog("ClientManager", "AddClient",
			"ClientDal_Add", err.Error())
		return  false, err.Error()
	}
	return  true, ""
}

func (c *clientManager)UpdateByClientId(clientId string, data *model.ClientDataModel)(success bool,message string){

	defer (*c.Log).SendInfoLog("ClientManager", "UpdateByClientId",
		clientId, data.ProjectId)

	if err:= (*c.ClientDal).UpdateById(clientId, data); err != nil {
		(*c.Log).SendErrorLog("ClientManager", "UpdateByClientId",
			"ClientDal_UpdateById", err.Error())
		return  false, err.Error()
	}
	return  true, ""

}

func (c *clientManager) GetByClientId (clientId string)(data *model.ClientDataModel, success bool,message string){

	defer (*c.Log).SendInfoLog("ClientManager", "GetByClientId", clientId)

	var client, err = (*c.ClientDal).GetById(clientId)
	if err != nil {
		(*c.Log).SendErrorLog("ClientManager", "GetByClientId",
			"ClientDal_GetById", err.Error())
	return nil, false, err.Error()
	}
	return client, true, ""
}