package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"log"
)

type offerBehaviorManager struct {
	Parser *JsonParser.IJsonParser
	OfferBehaviorDal *abstract.IOfferBehaviorDal
}

func OfferBehaviorManagerConstructor() *offerBehaviorManager {
	return &offerBehaviorManager{Parser: &IoC.JsonParser,
		OfferBehaviorDal: &IoC.OfferBehaviorDal}
}

func (o *offerBehaviorManager)AddOfferBehaviorData(data *[]byte)(success bool,message string){
	m := model.OfferBehaviorModel{}
	if err := (*o.Parser).DecodeJson(data, &m); err != nil {
		log.Fatal("OfferBehaviorManager", "AddOfferBehaviorData",
			"byte array to OfferBehaviorModel", "Json Parser Decode Err: ", err.Error())
		return false, err.Error()
	}

	defer log.Print("OfferBehaviorManager", "AddOfferBehaviorData",
		m.ClientId, m.ProjectId)

	if err:= (*o.OfferBehaviorDal).Add(&m); err != nil {
		log.Fatal("OfferBehaviorManager", "AddOfferBehaviorData_Add",
			"OfferBehaviorDal_Add", err.Error())
		return  false, err.Error()
	}
	return  true, ""
}

