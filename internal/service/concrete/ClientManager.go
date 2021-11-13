package concrete

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	jsonparser "StorageWorkerService/pkg/jsonParser"
)

type ClientManager struct {
	Parser jsonparser.IJsonParser
	ClientDal abstract.IClientDal
}

func (c *ClientManager)AddClient(data *[]byte)(success bool,message string){

	client := model.ClientDataModel{}
	c.Parser.DecodeJson(data, &client)

	err:= c.ClientDal.Add(&client)
	if err != nil {
		return  false, err.Error()
	}
	return  true, ""
}

func (c *ClientManager)UpdateByClientId(clientId string, data *model.ClientDataModel)(success bool,message string){

	err:= c.ClientDal.UpdateById(clientId, data)
	if err != nil {
		return  false, err.Error()
	}
	return  true, ""

}

func (c *ClientManager) GetByClientId (clientId string)(data *model.ClientDataModel, success bool,message string){

	data, err:= c.ClientDal.GetById(clientId)
	if err != nil {
	return nil, false, err.Error()
	}
	return data, true, ""
}