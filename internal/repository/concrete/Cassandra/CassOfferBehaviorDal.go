package Cassandra

import (
	"StorageWorkerService/internal/model"
	cassandra "StorageWorkerService/pkg/database/Cassandra"
	"fmt"
	"github.com/gocql/gocql"
)

type cassOfferBehaviorDal struct {
	Client *gocql.Session
	Table  string
}

func NewOfferBehaviorDal(Table string) *cassOfferBehaviorDal {
	return &cassOfferBehaviorDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassOfferBehaviorDal) Add(data *model.OfferBehaviorModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO %s(id, client_id, project_id, customer_id, version, offer_id, date_time, isBuy_offer, status) VALUES(?,?,?,?,?,?,?,?,?)", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId, data.Version, data.OfferId, data.DateTime, data.IsBuyOffer, data.Status).Exec(); err != nil {
		return err
	}
	return nil
}
