package model

import "time"

type InventoryModel struct{
	ClientId string
	ProjectId string
	CustomerId string
	MinorMine float32
	ModerateMine float32
	PreciousMine float32
	Items []Item
	Skills             []Item
	TemporaryAbilities []Item
	CreatedAt          time.Time
}

type Item struct {
	ItemType string
	Count int
}

type Skill struct {
	SkillType string
	Count int
}

type TemporaryAbility struct{
	AbilityType string
	Count int
}
