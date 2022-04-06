package model

import "time"

type GameSessionModel struct {
	Id                int64
	ClientId          int64
	ProjectId         int64
	CustomerId        int64
	SessionTime       float32
	CreatedAt         time.Time
	SessionStartTime  time.Time
	SessionFinishTime time.Time
	Status            bool
}
