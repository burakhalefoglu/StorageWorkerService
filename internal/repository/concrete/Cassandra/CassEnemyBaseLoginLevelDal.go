package Cassandra

import (
	"StorageWorkerService/internal/model"
	cassandra "StorageWorkerService/pkg/database/Cassandra"
	"fmt"
	"github.com/gocql/gocql"
)

type cassEnemyBaseLoginLevelDal struct {
	Client *gocql.Session
	Table  string
}

func NewEnemyBaseLoginLevelDal(Table string) *cassEnemyBaseLoginLevelDal {
	return &cassEnemyBaseLoginLevelDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassEnemyBaseLoginLevelDal) Add(data *model.EnemyBaseLoginLevelModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO %s(id, client_id, project_id, customer_id, level_name, level_index, playing_time, average_scores, date_time, is_dead, total_power_usage, status) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId, data.LevelName, data.LevelIndex, data.PlayingTime, data.AverageScores, data.DateTime, data.IsDead, data.TotalPowerUsage, data.Status).Exec(); err != nil {
		return err
	}
	return nil
}
