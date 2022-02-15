package Cassandra

import (
	"StorageWorkerService/internal/model"
	cassandra "StorageWorkerService/pkg/database/Cassandra"
	"fmt"
	"github.com/gocql/gocql"
)

type cassManuelFlowDal struct {
	Client *gocql.Session
	Table  string
}

func NewManuelFlowDal(Table string) *cassManuelFlowDal {
	return &cassManuelFlowDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassManuelFlowDal) Add(data *model.ManuelFlowModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO %s(id, client_id, project_id, customer_id, difficulty_level, date_time, status) VALUES(?,?,?,?,?,?,?)", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId, data.DifficultyLevel, data.DateTime, data.Status).Exec(); err != nil {
		return err
	}
	return nil
}
