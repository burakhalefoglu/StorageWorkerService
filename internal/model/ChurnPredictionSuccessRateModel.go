package model

import "time"

type ChurnPredictionSuccessRateModel struct {
	Id        int64
	ProjectId int64
	Value     string
	CreatedAt time.Time
	Status    bool
}
