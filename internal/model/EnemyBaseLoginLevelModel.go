package model

import "time"

type EnemyBaseLoginLevelModel struct{
	ClientId string
	ProjectId string
	CustomerId string
	LevelName string
	LevelIndex int
	PlayingTime int
	AverageScores int
	IsDead int
	TotalPowerUsage int
	DateTime time.Time
}
