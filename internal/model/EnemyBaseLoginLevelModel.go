package model

import "time"

type EnemyBaseLoginLevelModel struct {
	Id              int64
	ClientId        int64
	ProjectId       int64
	CustomerId      int64
	LevelName       string
	LevelIndex      int
	PlayingTime     int
	AverageScores   int
	IsDead          int
	TotalPowerUsage int
	DateTime        time.Time
	Status          bool
}
