package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
)

type clientManager struct {
	Parser    *JsonParser.IJsonParser
	ClientDal *abstract.IClientDal
}

func ClientManagerConstructor() *clientManager {
	return &clientManager{Parser: &IoC.JsonParser,
		ClientDal: &IoC.ClientDal}
}

func (c *clientManager) AddClient(data *[]byte) (success bool, message string) {

	client := model.ClientDataModel{}
	if err := (*c.Parser).DecodeJson(data, &client); err != nil {
		clogger.Error(&map[string]interface{}{
			"Json Parser Decode Err: ": err,
		})
		return false, err.Error()
	}

	if err := (*c.ClientDal).Add(&client); err != nil {
		clogger.Error(&map[string]interface{}{
			"ClientDal_Add": err,
		})
		return false, err.Error()
	}

	defer clogger.Info(&map[string]interface{}{
		fmt.Sprintf("Client: %d", client.Id): "added",
	})

	return true, ""
}

func (c *clientManager) UpdateByClientId(clientId int64, data *model.ClientDataModel) (success bool, message string) {

	if err := (*c.ClientDal).UpdateById(clientId, data); err != nil {
		clogger.Error(&map[string]interface{}{
			"ClientDal_UpdateById": err,
		})
		return false, err.Error()
	}

	defer clogger.Info(&map[string]interface{}{
		fmt.Sprintf("Client: %d", clientId): "update",
	})

	return true, ""

}

func (c *clientManager) GetByClientId(clientId int64) (data *model.ClientDataModel, success bool, message string) {

	var client, err = (*c.ClientDal).GetById(clientId)
	if err != nil {
		clogger.Error(&map[string]interface{}{
			"ClientDal_GetById": err.Error(),
		})
		return nil, false, err.Error()
	}
	defer clogger.Info(&map[string]interface{}{
		fmt.Sprintf("Client: %d", clientId): "get",
	})
	return client, true, ""
}
