package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"StorageWorkerService/pkg/logger"
)

type gameSessionManager struct {
	Parser *JsonParser.IJsonParser
	GameSessionDal *abstract.IGameSessionDal
	Log *logger.ILog
}

func GameSessionManagerConstructor() *gameSessionManager {
	return &gameSessionManager{Parser: &IoC.JsonParser,
		GameSessionDal: &IoC.GameSessionDal,
		Log: &IoC.Logger}
}

func (g *gameSessionManager)AddGameSessionData(data *[]byte)(success bool,message string){

	m := model.GameSessionModel{}
	if err := (*g.Parser).DecodeJson(data, &m); err != nil {
		(*g.Log).SendErrorLog("GameSessionManager", "AddGameSessionData",
			"byte array to GameSessionModel", "Json Parser Decode Err: ", err.Error())
		return false, err.Error()
	}

	defer (*g.Log).SendInfoLog("GameSessionManager", "AddGameSessionData",
		m.ClientId, m.ProjectId)

	if err:= (*g.GameSessionDal).Add(&m); err != nil {
		(*g.Log).SendErrorLog("GameSessionManager", "AddGameSessionData",
			"GameSessionDal_Add", err.Error())
		return  false, err.Error()
	}
	return  true, ""
}