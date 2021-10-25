package model

import "time"

type ScreenSwipeModel struct{
	ClientId string
	ProjectId string
	CustomerId string
	SwipeDirection int
	StartLocX float32
	StartLocY float32
	FinishLocX float32
	FinishLocY float32
	LevelName string
	LevelIndex int
	CreatedAt time.Time
}
