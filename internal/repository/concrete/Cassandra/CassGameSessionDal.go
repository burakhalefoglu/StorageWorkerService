package Cassandra

import (
	"StorageWorkerService/internal/model"
	cassandra "StorageWorkerService/pkg/database/Cassandra"
	"fmt"
	"github.com/gocql/gocql"
)

type cassGameSessionDal struct {
	Client *gocql.Session
	Table  string
}

func NewGameSessionDal(Table string) *cassGameSessionDal {
	return &cassGameSessionDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassGameSessionDal) Add(data *model.GameSessionModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO %s(id, client_id, project_id, customer_id, session_start_time, session_finish_time, session_time_minute, status) VALUES(?,?,?,?,?,?,?,?)", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId, data.SessionStartTime, data.SessionFinishTime, data.SessionTimeMinute, data.Status).Exec(); err != nil {
		return err
	}
	return nil
}
