package model

import "time"

type ClientDataModel struct{

	ClientId string
	ProjectId string
	IsPaidClient int
	CreatedAt time.Time
	PaidTime time.Time
}