package model

import "time"

type LevelBaseSessionModel struct {
	Id                int64
	ClientId          int64
	ProjectId         int64
	CustomerId        int64
	LevelName         string
	LevelIndex        int
	SessionTimeMinute float32
	SessionStartTime  time.Time
	SessionFinishTime time.Time
	Status            bool
}
