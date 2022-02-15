package model

import "time"

type InventoryModelDto struct {
	Id                 int64
	ClientId           int64
	ProjectId          int64
	CustomerId         int64
	MinorMine          float32
	ModerateMine       float32
	PreciousMine       float32
	Items              []ItemModel
	Skills             []SkillModel
	TemporaryAbilities []TemporaryAbilityModel
	CreatedAt          time.Time
	Status             bool
}

type InventoryModel struct {
	Id           int64
	ClientId     int64
	ProjectId    int64
	CustomerId   int64
	MinorMine    float32
	ModerateMine float32
	PreciousMine float32
	CreatedAt    time.Time
	Status       bool
}

type ItemModel struct {
	Id          int64
	InventoryId int64
	ItemType    string
	Count       int
}

type SkillModel struct {
	Id          int64
	InventoryId int64
	SkillType   string
	Count       int
}

type TemporaryAbilityModel struct {
	Id          int64
	InventoryId int64
	AbilityType string
	Count       int
}
