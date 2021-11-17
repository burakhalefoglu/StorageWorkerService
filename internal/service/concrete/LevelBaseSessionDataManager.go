package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"StorageWorkerService/pkg/logger"
)

type levelBaseSessionManager struct {
	Parser *JsonParser.IJsonParser
	LevelBaseSessionDal *abstract.ILevelBaseSessionDal
	Log *logger.ILog
}

func LevelBaseSessionManagerConstructor() *levelBaseSessionManager {
	return &levelBaseSessionManager{Parser: &IoC.JsonParser,
		LevelBaseSessionDal: &IoC.LevelBaseSessionDal,
		Log: &IoC.Logger}
}

func (lvl *levelBaseSessionManager)AddLevelBaseSessionData(data *[]byte)(success bool,message string){

	m := model.LevelBaseSessionModel{}
	if err := (*lvl.Parser).DecodeJson(data, &m); err != nil {
		(*lvl.Log).SendErrorLog("LevelBaseSessionManager", "AddLevelBaseSessionData",
			"byte array to LevelBaseSessionModel", "Json Parser Decode Err: ", err.Error())
		return false, err.Error()
	}

	defer (*lvl.Log).SendInfoLog("LevelBaseSessionManager", "AddLevelBaseSessionData",
		m.ClientId, m.ProjectId)

	if err := (*lvl.LevelBaseSessionDal).Add(&m); err != nil {
		(*lvl.Log).SendErrorLog("LevelBaseSessionManager", "AddLevelBaseSessionData",
			"LevelBaseSessionDal_Add", err.Error())
		return  false, err.Error()
	}
	return  true, ""
}