package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
)

type screenClickManager struct {
	Parser         *JsonParser.IJsonParser
	ScreenClickDal *abstract.IScreenClickDal
}

func ScreenClickManagerConstructor() *screenClickManager {
	return &screenClickManager{Parser: &IoC.JsonParser,
		ScreenClickDal: &IoC.ScreenClickDal}
}

func (scr *screenClickManager) AddScreenClickData(data *[]byte) (success bool, message string) {

	screenClickData := model.ScreenClickModel{}
	if err := (*scr.Parser).DecodeJson(data, &screenClickData); err != nil {
		clogger.Error(&map[string]interface{}{
			"Json Parser Decode Err: ": err,
		})
		return false, err.Error()
	}

	defer clogger.Info(&map[string]interface{}{
		fmt.Sprintf("Data: %d", screenClickData.ClientId): "added",
	})

	if err := (*scr.ScreenClickDal).Add(&screenClickData); err != nil {
		clogger.Error(&map[string]interface{}{
			"Repository_Add Error": err,
		})
		return false, err.Error()
	}
	return true, ""
}
