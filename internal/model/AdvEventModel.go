package model

import "time"

type AdvEventDataModel struct {
	Id            int64
	ClientId      int64
	ProjectId     int64
	CustomerId    int64
	LevelName     string
	LevelIndex    int32
	AdvType       string
	InMinutes     float32
	TriggeredTime time.Time
	Status        bool
}
