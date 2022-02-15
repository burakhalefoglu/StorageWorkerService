package Cassandra

import (
	"StorageWorkerService/internal/model"
	cassandra "StorageWorkerService/pkg/database/Cassandra"
	"fmt"
	"github.com/gocql/gocql"
)

type cassLocationDal struct {
	Client *gocql.Session
	Table  string
}

func NewLocationDal(Table string) *cassLocationDal {
	return &cassLocationDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassLocationDal) Add(data *model.LocationModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO %s(id, client_id, project_id, customer_id, continent, country, city, query, region, org, created_at, status) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId, data.Continent, data.Continent, data.Country, data.City, data.Query, data.Region, data.Org, data.CreatedAt, data.Status).Exec(); err != nil {
		return err
	}
	return nil
}
