package Cassandra

import (
	"StorageWorkerService/internal/model"
	cassandra "StorageWorkerService/pkg/database/Cassandra"
	"fmt"
	"github.com/gocql/gocql"
)

type cassLevelBaseSessionDal struct {
	Client *gocql.Session
	Table  string
}

func NewLevelBaseSessionDal(Table string) *cassLevelBaseSessionDal {
	return &cassLevelBaseSessionDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassLevelBaseSessionDal) Add(data *model.LevelBaseSessionModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO %s(id, client_id, project_id, customer_id, level_name, level_index, session_time_minute, session_start_time, session_finish_time, status) VALUES(?,?,?,?,?,?,?,?,?,?)", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId, data.LevelName, data.LevelIndex, data.SessionTimeMinute, data.SessionStartTime, data.SessionFinishTime, data.Status).Exec(); err != nil {
		return err
	}
	return nil
}
