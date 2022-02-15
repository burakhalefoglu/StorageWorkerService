package model

import "time"

type EnemyBaseLevelFailModel struct {
	Id                         int64
	ClientId                   int64
	ProjectId                  int64
	CustomerId                 int64
	LevelName                  string
	LevelIndex                 int
	FailTimeAfterLevelStarting int
	FailLocationX              float32
	FailLocationY              float32
	FailLocationZ              float32
	DateTime                   time.Time
	Status                     bool
}
