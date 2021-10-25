package model

import "time"

type OfferBehaviorModel struct{
	ClientId string
	ProjectId string
	CustomerId string
	OfferId int
	IsBuyOffer int
	DateTime time.Time
}
