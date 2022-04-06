package model

import "time"

type AdvStrategyBehaviorModel struct {
	Id         int64
	StrategyId int64
	CreatedAt  time.Time
	ClientId   int64
	ProjectId  int64
	Version    int32
	Name       string
	Status     bool
}
