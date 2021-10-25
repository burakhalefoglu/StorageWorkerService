package model

import "time"

type ScreenClickModel struct{

	ClientId string
	ProjectId string
	CustomerId string
	StartLocX float32
	StartLocY float32
	FinishLocX float32
	FinishLocY float32
	LevelName string
	LevelIndex int
	TabCount int
	FingerID int
	CreatedAt time.Time
}
