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
	parseErr := (*g.Parser).DecodeJson(data, &m)
	if parseErr != nil {
		(*g.Log).SendErrorLog("GameSessionManager", "AddGameSessionData",
			"DecodeJson Error", parseErr.Error())
		return false, parseErr.Error()
	}

	defer (*g.Log).SendInfoLog("GameSessionManager", "AddGameSessionData",
		m.ClientId, m.ProjectId)

	err:= (*g.GameSessionDal).Add(&m)
	if err != nil {
		(*g.Log).SendErrorLog("GameSessionManager", "GameSessionManager_Add",
			m.ClientId, m.ProjectId, err.Error())
		return  false, err.Error()
	}
	return  true, ""
}