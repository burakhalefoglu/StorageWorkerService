package model

import "time"

type ChurnBlockerMlResultModel struct{
	ClientId string
	ProjectId string
	CustomerId string
	ModelType string
	ModelResult float32
	DateTime time.Time
}

