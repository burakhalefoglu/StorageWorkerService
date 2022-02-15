package Cassandra

import (
	"StorageWorkerService/internal/model"
	cassandra "StorageWorkerService/pkg/database/Cassandra"
	"fmt"
	"github.com/gocql/gocql"
)

type cassEnemyBaseLevelFailDal struct {
	Client *gocql.Session
	Table  string
}

func NewEnemyBaseLevelFailDal(Table string) *cassEnemyBaseLevelFailDal {
	return &cassEnemyBaseLevelFailDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassEnemyBaseLevelFailDal) Add(data *model.EnemyBaseLevelFailModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO %s(id, client_id, project_id, customer_id, fail_time_after_level_starting, level_name, level_index, fail_location_x, fail_location_y, fail_location_z, date_time,  status) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId, data.FailTimeAfterLevelStarting, data.LevelName, data.LevelIndex, data.FailLocationX, data.FailLocationY, data.FailLocationZ, data.DateTime, data.Status).Exec(); err != nil {
		return err
	}
	return nil
}
