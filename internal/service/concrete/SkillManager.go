package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"log"
)

type skillManager struct {
	Parser   *JsonParser.IJsonParser
	SkillDal *abstract.ISkillDal
}

func NewSkillManager() *skillManager {
	return &skillManager{Parser: &IoC.JsonParser,
		SkillDal: &IoC.SkillDal}
}

func (i *skillManager) AddSkillData(data *model.SkillModel) (success bool, message string) {

	defer log.Print("skillManager", "AddSkillData",
		data.Id)

	if err := (*i.SkillDal).Add(data); err != nil {
		log.Fatal("skillManager", "AddSkillData",
			"SkillDal_Add", err.Error())
		return false, err.Error()
	}
	return true, ""
}
