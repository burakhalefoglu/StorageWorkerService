package model

import "time"

type ScreenClickModel struct {
	Id         int64
	ClientId   int64
	ProjectId  int64
	CustomerId int64
	StartLocX  float32
	StartLocY  float32
	FinishLocX float32
	FinishLocY float32
	LevelName  string
	LevelIndex int
	TabCount   int
	FingerId   int
	CreatedAt  time.Time
	Status     bool
}
