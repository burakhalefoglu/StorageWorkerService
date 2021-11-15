package concrete

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	jsonparser "StorageWorkerService/pkg/jsonParser"
)

type offerBehaviorManager struct {
	Parser jsonparser.IJsonParser
	OfferBehaviorDal abstract.IOfferBehaviorDal
}

func OfferBehaviorManagerConstructor(parser jsonparser.IJsonParser,
	offerBehaviorDal abstract.IOfferBehaviorDal) *offerBehaviorManager {
	return &offerBehaviorManager{Parser: parser, OfferBehaviorDal: offerBehaviorDal}
}

func (o *offerBehaviorManager)AddOfferBehaviorData(data *[]byte)(success bool,message string){
	m := model.OfferBehaviorModel{}
	o.Parser.DecodeJson(data, &m)

	err:= o.OfferBehaviorDal.Add(&m)
	if err != nil {
		return  false, err.Error()
	}
	return  true, ""
}

