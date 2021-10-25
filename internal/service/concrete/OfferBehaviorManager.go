package concrete

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	jsonparser "StorageWorkerService/pkg/jsonParser"
)

type OfferBehaviorManager struct {
	Parser jsonparser.IJsonParser
	OfferBehaviorDal abstract.IOfferBehaviorDal
}


func (o *OfferBehaviorManager)AddOfferBehaviorData(data *[]byte)(success bool,message string){
	m := model.OfferBehaviorModel{}
	o.Parser.DecodeJson(data, &m)

	err:= o.OfferBehaviorDal.Add(&m)
	if err != nil {
		return  false, err.Error()
	}
	return  true, ""
}

