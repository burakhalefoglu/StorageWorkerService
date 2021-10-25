package model

import "time"

type LevelBaseSessionModel struct{

	ClientId string
	ProjectId string
	CustomerId string
	LevelName string
	LevelIndex int
	SessionTimeMinute float32
	SessionStartTime time.Time
	SessionFinishTime time.Time
}
