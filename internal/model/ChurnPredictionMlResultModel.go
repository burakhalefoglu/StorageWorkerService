package model

import "time"

type ChurnPredictionMlResultModel struct {
	Id          int64
	ClientId    int64
	ProjectId   int64
	CustomerId  int64
	ModelType   string
	ModelResult float32
	DateTime    time.Time
	Status      bool
}
