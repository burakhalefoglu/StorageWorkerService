package concrete

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	jsonparser "StorageWorkerService/pkg/jsonParser"
)

type levelBaseSessionManager struct {
	Parser jsonparser.IJsonParser
	LevelBaseSessionDal abstract.ILevelBaseSessionDal
}

func LevelBaseSessionManagerConstructor(parser jsonparser.IJsonParser,
	levelBaseSessionDal abstract.ILevelBaseSessionDal) *levelBaseSessionManager {
	return &levelBaseSessionManager{Parser: parser, LevelBaseSessionDal: levelBaseSessionDal}
}

func (lvl *levelBaseSessionManager)AddLevelBaseSessionData(data *[]byte)(success bool,message string){

	m := model.LevelBaseSessionModel{}
	lvl.Parser.DecodeJson(data, &m)

	err:= lvl.LevelBaseSessionDal.Add(&m)
	if err != nil {
		return  false, err.Error()
	}
	return  true, ""
}