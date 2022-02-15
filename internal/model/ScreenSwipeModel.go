package model

import "time"

type ScreenSwipeModel struct {
	Id             int64
	ClientId       int64
	ProjectId      int64
	CustomerId     int64
	StartLocX      float32
	StartLocY      float32
	FinishLocX     float32
	FinishLocY     float32
	SwipeDirection int
	LevelName      string
	LevelIndex     int
	CreatedAt      time.Time
	Status         bool
}
