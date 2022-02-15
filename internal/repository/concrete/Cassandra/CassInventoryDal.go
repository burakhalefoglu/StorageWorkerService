package Cassandra

import (
	"StorageWorkerService/internal/model"
	cassandra "StorageWorkerService/pkg/database/Cassandra"
	"fmt"
	"github.com/gocql/gocql"
)

type cassInventoryDal struct {
	Client *gocql.Session
	Table  string
}

func NewInventoryDal(Table string) *cassInventoryDal {
	return &cassInventoryDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassInventoryDal) Add(data *model.InventoryModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO %s(id, client_id, project_id, customer_id, minor_mine, moderate_mine, precious_mine, created_at, status) VALUES(?,?,?,?,?,?,?,?,?)", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId, data.MinorMine, data.ModerateMine, data.PreciousMine, data.CreatedAt, data.Status).Exec(); err != nil {
		return err
	}
	return nil
}
