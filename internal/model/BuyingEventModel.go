package model

import "time"

type BuyingEventModel struct{
	ClientId string
	ProjectId string
	CustomerId string
	LevelName string
	LevelIndex int32
	ProductType string
	InWhatMinutes float32
	TriggeredTime time.Time
}
