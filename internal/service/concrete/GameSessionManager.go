package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
)

type gameSessionManager struct {
	Parser         *JsonParser.IJsonParser
	GameSessionDal *abstract.IGameSessionDal
}

func GameSessionManagerConstructor() *gameSessionManager {
	return &gameSessionManager{Parser: &IoC.JsonParser,
		GameSessionDal: &IoC.GameSessionDal}
}

func (g *gameSessionManager) AddGameSessionData(data *[]byte) (success bool, message string) {

	m := model.GameSessionModel{}
	if err := (*g.Parser).DecodeJson(data, &m); err != nil {
		clogger.Error(&map[string]interface{}{
			"Json Parser Decode Err: ": err,
		})
		return false, err.Error()
	}

	defer clogger.Info(&map[string]interface{}{
		fmt.Sprintf("Data: %d", m.ClientId): "added",
	})

	if err := (*g.GameSessionDal).Add(&m); err != nil {
		clogger.Error(&map[string]interface{}{
			"Repository_Add Error": err,
		})
		return false, err.Error()
	}
	return true, ""
}
