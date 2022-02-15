package Cassandra

import (
	"StorageWorkerService/internal/model"
	cassandra "StorageWorkerService/pkg/database/Cassandra"
	"fmt"
	"github.com/gocql/gocql"
)

type cassBuyingEventDal struct {
	Client *gocql.Session
	Table  string
}

func NewCassBuyingEventDal(Table string) *cassBuyingEventDal {
	return &cassBuyingEventDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassBuyingEventDal) Add(data *model.BuyingEventModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO %s(id, client_id, project_id, customer_id, level_name, level_index, product_type, in_minutes, triggered_time, status) VALUES(?,?,?,?,?,?,?,?,?,?)", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId, data.LevelName, data.LevelIndex, data.ProductType, data.InMinutes, data.TriggeredTime, data.Status).Exec(); err != nil {
		return err
	}
	return nil
}
