package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
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

	defer clogger.Info(&map[string]interface{}{
		fmt.Sprintf("Data: %d", data.Id): "added",
	})

	if err := (*i.SkillDal).Add(data); err != nil {
		clogger.Error(&map[string]interface{}{
			"Repository_Add Error": err,
		})
		return false, err.Error()
	}
	return true, ""
}
