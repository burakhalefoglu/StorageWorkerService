package model

import "time"

type EnemyBaseLevelFailModel struct{

	ClientId string
	ProjectId string
	CustomerId string
	LevelName string
	LevelIndex int
	DiyingTimeAfterLevelStarting int
	FailLocationX float32
	FailLocationY float32
	FailLocationZ float32
	DateTime time.Time
}
