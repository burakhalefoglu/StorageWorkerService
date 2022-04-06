package concrete

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/repository/abstract"
	service "StorageWorkerService/internal/service/abstract"
	JsonParser "StorageWorkerService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
)

type inventoryManager struct {
	Parser                  *JsonParser.IJsonParser
	InventoryDal            *abstract.IInventoryDal
	SkillService            *service.ISkillService
	ItemService             *service.IItemService
	temporaryAbilityService *service.ITemporaryAbilityService
}

func NewInventoryManager() *inventoryManager {
	return &inventoryManager{Parser: &IoC.JsonParser,
		InventoryDal:            &IoC.InventoryDal,
		ItemService:             &IoC.ItemService,
		SkillService:            &IoC.SkillService,
		temporaryAbilityService: &IoC.TemporaryAbilityService}
}

func (i *inventoryManager) AddInventoryData(data *[]byte) (success bool, message string) {

	mDto := model.InventoryModelDto{}
	if err := (*i.Parser).DecodeJson(data, &mDto); err != nil {
		clogger.Error(&map[string]interface{}{
			"Json Parser Decode Err: ": err,
		})
		return false, err.Error()
	}

	defer clogger.Info(&map[string]interface{}{
		fmt.Sprintf("Data: %d", mDto.ClientId): "added",
	})

	m := model.InventoryModel{
		Id:           mDto.Id,
		ClientId:     mDto.ClientId,
		ProjectId:    mDto.ProjectId,
		CustomerId:   mDto.CustomerId,
		MinorMine:    mDto.MinorMine,
		ModerateMine: mDto.ModerateMine,
		PreciousMine: mDto.PreciousMine,
		CreatedAt:    mDto.CreatedAt,
		Status:       mDto.Status,
	}

	if err := (*i.InventoryDal).Add(&m); err != nil {
		clogger.Error(&map[string]interface{}{
			"InventoryDal_Add: ": err,
		})
		return false, err.Error()
	}

	for _, item := range mDto.Items {
		if s, m := (*i.ItemService).AddItemData(&item); s == false {
			clogger.Error(&map[string]interface{}{
				"err message: ": m,
			})
		}
	}

	for _, skill := range mDto.Skills {
		if s, m := (*i.SkillService).AddSkillData(&skill); s == false {
			clogger.Error(&map[string]interface{}{
				"err message: ": m,
			})
		}
	}

	for _, temporaryAbility := range mDto.TemporaryAbilities {
		if s, m := (*i.temporaryAbilityService).AddTemporaryAbilityData(&temporaryAbility); s == false {
			clogger.Error(&map[string]interface{}{
				"err message: ": m,
			})
		}
	}

	return true, ""
}
