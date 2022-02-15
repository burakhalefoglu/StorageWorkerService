package abstract

import "StorageWorkerService/internal/model"

type IInventoryService interface {
	AddInventoryData(data *[]byte) (success bool, message string)
}

type IItemService interface {
	AddItemData(data *model.ItemModel) (success bool, message string)
}

type ISkillService interface {
	AddSkillData(data *model.SkillModel) (success bool, message string)
}

type ITemporaryAbilityService interface {
	AddTemporaryAbilityData(data *model.TemporaryAbilityModel) (success bool, message string)
}
