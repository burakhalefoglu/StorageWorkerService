package model

import "time"

type ManuelFlowModel struct{
	ClientId string
	ProjectId string
	CustomerId string
	DifficultyLevel int
	DateTime time.Time
}
