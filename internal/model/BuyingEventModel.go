package model

import "time"

type BuyingEventModel struct {
	Id            int64
	ClientId      int64
	ProjectId     int64
	CustomerId    int64
	LevelName     string
	LevelIndex    int32
	ProductType   string
	InMinutes     float32
	TriggeredTime time.Time
	Status        bool
}
