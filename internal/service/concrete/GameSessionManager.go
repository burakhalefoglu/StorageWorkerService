package concrete

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	jsonparser "StorageWorkerService/pkg/jsonParser"
)

type GameSessionManager struct {
	Parser jsonparser.IJsonParser
	GameSessionDal abstract.IGameSessionDal
}

func (g *GameSessionManager)AddGameSessionData(data *[]byte)(success bool,message string){

	m := model.GameSessionModel{}
	g.Parser.DecodeJson(data, &m)

	err:= g.GameSessionDal.Add(&m)
	if err != nil {
		return  false, err.Error()
	}
	return  true, ""
}