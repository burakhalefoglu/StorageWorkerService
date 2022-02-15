package Cassandra

import (
	"StorageWorkerService/internal/model"
	cassandra "StorageWorkerService/pkg/database/Cassandra"
	"fmt"
	"github.com/gocql/gocql"
)

type cassAdvEventDal struct {
	Client *gocql.Session
	Table  string
}

func NewCassAdvEventDal(Table string) *cassAdvEventDal {
	return &cassAdvEventDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassAdvEventDal) Add(data *model.AdvEventDataModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO %s(id, client_id, project_id, customer_id, level_name, level_index, adv_type, in_minutes, trigger_time, status) VALUES(?,?,?,?,?,?,?,?,?,?)", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId, data.LevelName, data.LevelIndex, data.AdvType, data.InMinutes, data.TriggeredTime, data.Status).Exec(); err != nil {
		return err
	}
	return nil
}
