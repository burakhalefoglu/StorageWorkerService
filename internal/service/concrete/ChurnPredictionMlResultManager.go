package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"log"
)

type churnPredictionMlResultManager struct {
	Parser *JsonParser.IJsonParser
	ChurnPredictionMlResultDal *abstract.IChurnPredictionMlResultDal
}

func ChurnPredictionMlResultManagerConstructor() *churnPredictionMlResultManager {
	return &churnPredictionMlResultManager{Parser: &IoC.JsonParser,
		ChurnPredictionMlResultDal: &IoC.ChurnPredictionMlResultDal}
}

func (c *churnPredictionMlResultManager) AddChurnPredictionMlResultData(data *[]byte)(success bool,message string){

	churnPredictionModel := model.ChurnPredictionMlResultModel{}
	if err := (*c.Parser).DecodeJson(data, &churnPredictionModel); err != nil {
		log.Fatal("ChurnPredictionMlResultManager", "AddChurnPredictionMlResultData",
			"byte array to ChurnPredictionMlResultModel", "Json Parser Decode Err: ", err.Error())
		return false, err.Error()
	}

	defer log.Print("ChurnPredictionMlResultManager", "AddChurnPredictionMlResultData",
		churnPredictionModel.ClientId, churnPredictionModel.ProjectId)

	if err:= (*c.ChurnPredictionMlResultDal).Add(&churnPredictionModel); err != nil {
		log.Fatal("ChurnPredictionMlResultManager", "AddChurnPredictionMlResultData",
			"ChurnPredictionMlResultDal_Add", err.Error())
		return  false, err.Error()
	}
	return  true, ""
}