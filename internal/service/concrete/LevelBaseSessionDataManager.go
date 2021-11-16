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
	parseErr := (*lvl.Parser).DecodeJson(data, &m)
	if parseErr != nil {
		(*lvl.Log).SendErrorLog("LevelBaseSessionManager", "AddLevelBaseSessionData",
			"DecodeJson Error", parseErr.Error())
		return false, parseErr.Error()
	}

	defer (*lvl.Log).SendInfoLog("LevelBaseSessionManager", "AddLevelBaseSessionData",
		m.ClientId, m.ProjectId)

	err:= (*lvl.LevelBaseSessionDal).Add(&m)
	if err != nil {
		(*lvl.Log).SendErrorLog("LevelBaseSessionManager", "AddLevelBaseSessionData_Add",
			m.ClientId, m.ProjectId, err.Error())
		return  false, err.Error()
	}
	return  true, ""
}