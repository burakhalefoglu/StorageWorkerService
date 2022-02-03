package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"log"
)

type levelBaseSessionManager struct {
	Parser *JsonParser.IJsonParser
	LevelBaseSessionDal *abstract.ILevelBaseSessionDal
}

func LevelBaseSessionManagerConstructor() *levelBaseSessionManager {
	return &levelBaseSessionManager{Parser: &IoC.JsonParser,
		LevelBaseSessionDal: &IoC.LevelBaseSessionDal}
}

func (lvl *levelBaseSessionManager)AddLevelBaseSessionData(data *[]byte)(success bool,message string){

	m := model.LevelBaseSessionModel{}
	if err := (*lvl.Parser).DecodeJson(data, &m); err != nil {
		log.Fatal("LevelBaseSessionManager", "AddLevelBaseSessionData",
			"byte array to LevelBaseSessionModel", "Json Parser Decode Err: ", err.Error())
		return false, err.Error()
	}

	defer log.Print("LevelBaseSessionManager", "AddLevelBaseSessionData",
		m.ClientId, m.ProjectId)

	if err := (*lvl.LevelBaseSessionDal).Add(&m); err != nil {
		log.Fatal("LevelBaseSessionManager", "AddLevelBaseSessionData",
			"LevelBaseSessionDal_Add", err.Error())
		return  false, err.Error()
	}
	return  true, ""
}