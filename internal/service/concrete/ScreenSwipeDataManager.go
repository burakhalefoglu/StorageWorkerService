package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
)

type screenSwipeManager struct {
	Parser         *JsonParser.IJsonParser
	ScreenSwipeDal *abstract.IScreenSwipeDal
}

func ScreenSwipeManagerConstructor() *screenSwipeManager {
	return &screenSwipeManager{Parser: &IoC.JsonParser,
		ScreenSwipeDal: &IoC.ScreenSwipeDal}
}

func (scr *screenSwipeManager) AddScreenSwipeData(data *[]byte) (success bool, message string) {

	screenSwipeData := model.ScreenSwipeModel{}
	if err := (*scr.Parser).DecodeJson(data, &screenSwipeData); err != nil {
		clogger.Error(&map[string]interface{}{
			"Json Parser Decode Err: ": err,
		})
		return false, err.Error()
	}

	defer clogger.Info(&map[string]interface{}{
		fmt.Sprintf("Data: %d", screenSwipeData.ClientId): "added",
	})

	if err := (*scr.ScreenSwipeDal).Add(&screenSwipeData); err != nil {
		clogger.Error(&map[string]interface{}{
			"Repository_Add Error": err,
		})
		return false, err.Error()
	}
	return true, ""
}
