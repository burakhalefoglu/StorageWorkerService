package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"log"
)

type manuelFlowManager struct {
	Parser *JsonParser.IJsonParser
	ManuelFlowDal *abstract.IManuelFlowDal
}

func ManuelFlowManagerConstructor() *manuelFlowManager {
	return &manuelFlowManager{Parser: &IoC.JsonParser,
		ManuelFlowDal: &IoC.ManuelFlowDal}
}

func (f *manuelFlowManager)AddManuelFlowData(data *[]byte)(success bool,message string){
	m := model.ManuelFlowModel{}
	if err := (*f.Parser).DecodeJson(data, &m); err != nil {
		log.Fatal("ManuelFlowManager", "AddManuelFlowData",
			"byte array to ManuelFlowModel", "Json Parser Decode Err: ", err.Error())
		return false, err.Error()
	}
	defer log.Print("ManuelFlowManager", "AddManuelFlowData",
		m.ClientId, m.ProjectId)

	if err:= (*f.ManuelFlowDal).Add(&m); err != nil {
		log.Fatal("ManuelFlowManager", "AddManuelFlowData",
			"ManuelFlowDal_Add", err.Error())
		return  false, err.Error()
	}
	return  true, ""
}