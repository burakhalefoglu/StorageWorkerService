package model

import "time"

type LocationModel struct {
	Id         int64
	ClientId   int64
	ProjectId  int64
	CustomerId int64
	Continent  string
	Country    string
	City       string
	Query      string
	Region     string
	Org        string
	CreatedAt  time.Time
	Status     bool
}
