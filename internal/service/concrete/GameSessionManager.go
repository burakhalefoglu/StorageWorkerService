package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"log"
)

type gameSessionManager struct {
	Parser *JsonParser.IJsonParser
	GameSessionDal *abstract.IGameSessionDal
}

func GameSessionManagerConstructor() *gameSessionManager {
	return &gameSessionManager{Parser: &IoC.JsonParser,
		GameSessionDal: &IoC.GameSessionDal}
}

func (g *gameSessionManager)AddGameSessionData(data *[]byte)(success bool,message string){

	m := model.GameSessionModel{}
	if err := (*g.Parser).DecodeJson(data, &m); err != nil {
		log.Fatal("GameSessionManager", "AddGameSessionData",
			"byte array to GameSessionModel", "Json Parser Decode Err: ", err.Error())
		return false, err.Error()
	}

	defer log.Print("GameSessionManager", "AddGameSessionData",
		m.ClientId, m.ProjectId)

	if err:= (*g.GameSessionDal).Add(&m); err != nil {
		log.Fatal("GameSessionManager", "AddGameSessionData",
			"GameSessionDal_Add", err.Error())
		return  false, err.Error()
	}
	return  true, ""
}