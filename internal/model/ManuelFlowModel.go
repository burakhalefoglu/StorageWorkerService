package model

import "time"

type ManuelFlowModel struct {
	Id              int64
	ClientId        int64
	ProjectId       int64
	CustomerId      int64
	DifficultyLevel int
	DateTime        time.Time
	Status          bool
}
