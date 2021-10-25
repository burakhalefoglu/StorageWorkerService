package concrete

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	jsonparser "StorageWorkerService/pkg/jsonParser"
)

type LevelBaseSessionManager struct {
	Parser jsonparser.IJsonParser
	LevelBaseSessionDal abstract.ILevelBaseSessionDal
}

func (lvl *LevelBaseSessionManager)AddLevelBaseSessionData(data *[]byte)(success bool,message string){

	m := model.LevelBaseSessionModel{}
	lvl.Parser.DecodeJson(data, &m)

	err:= lvl.LevelBaseSessionDal.Add(&m)
	if err != nil {
		return  false, err.Error()
	}
	return  true, ""
}