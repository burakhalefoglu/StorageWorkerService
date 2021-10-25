package model

import "time"

type AdvEventDataModel struct{
	ClientId string
	ProjectId string
	CustomerId string
	LevelName string
	LevelIndex int32
	AdvType string
	InMinutes     float32
	TriggeredTime time.Time
}
