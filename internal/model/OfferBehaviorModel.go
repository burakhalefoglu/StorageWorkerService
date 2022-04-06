package model

import "time"

type OfferBehaviorModel struct {
	Id         int64
	ClientId   int64
	ProjectId  int64
	CustomerId int64
	Version    int16
	OfferId    int
	IsBuyOffer int8
	CreatedAt  time.Time
	Status     bool
}
