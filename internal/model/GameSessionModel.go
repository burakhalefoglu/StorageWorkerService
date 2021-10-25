package model

import "time"

type GameSessionModel struct{

	ClientId string
	ProjectId string
	CustomerId string
	SessionTimeMinute float32
	SessionStartTime time.Time
	SessionFinishTime time.Time
}
