package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
)

type levelBaseSessionManager struct {
	Parser              *JsonParser.IJsonParser
	LevelBaseSessionDal *abstract.ILevelBaseSessionDal
}

func LevelBaseSessionManagerConstructor() *levelBaseSessionManager {
	return &levelBaseSessionManager{Parser: &IoC.JsonParser,
		LevelBaseSessionDal: &IoC.LevelBaseSessionDal}
}

func (lvl *levelBaseSessionManager) AddLevelBaseSessionData(data *[]byte) (success bool, message string) {

	m := model.LevelBaseSessionModel{}
	if err := (*lvl.Parser).DecodeJson(data, &m); err != nil {
		clogger.Error(&map[string]interface{}{
			"Json Parser Decode Err: ": err,
		})
		return false, err.Error()
	}

	defer clogger.Info(&map[string]interface{}{
		fmt.Sprintf("Data: %d", m.ClientId): "added",
	})

	if err := (*lvl.LevelBaseSessionDal).Add(&m); err != nil {
		clogger.Error(&map[string]interface{}{
			"Repository_Add Error": err,
		})
		return false, err.Error()
	}
	return true, ""
}
