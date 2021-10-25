package model

import "time"

type LocationModel struct{

	ClientId string
	ProjectId string
	CustomerId string
	Continent string
	Country string
	City string
	Query string
	Region string
	Org string
	CreatedAt time.Time
}
