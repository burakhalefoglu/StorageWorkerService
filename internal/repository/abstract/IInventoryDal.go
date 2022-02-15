package abstract

import (
	"StorageWorkerService/internal/model"
)

type IInventoryDal interface {
	Add(model *model.InventoryModel) error
}

type IItemDal interface {
	Add(model *model.ItemModel) error
}

type ISkillDal interface {
	Add(model *model.SkillModel) error
}

type ITemporaryAbilityDal interface {
	Add(model *model.TemporaryAbilityModel) error
}
