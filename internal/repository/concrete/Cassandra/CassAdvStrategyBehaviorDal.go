package Cassandra

import (
	"StorageWorkerService/internal/model"
	cassandra "StorageWorkerService/pkg/database/Cassandra"
	"fmt"

	"github.com/gocql/gocql"
)

type cassAdvStrategyBehaviorDal struct {
	Client *gocql.Session
	Table  string
}

func NewCassAdvStrategyBehaviorDal(Table string) *cassAdvStrategyBehaviorDal {
	return &cassAdvStrategyBehaviorDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassAdvStrategyBehaviorDal) Add(data *model.AdvStrategyBehaviorModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO %s(id, project_id, client_id, strategy_id, name, version, created_at, status) VALUES(?,?,?,?,?,?,?,?)", m.Table),
		data.Id, data.ProjectId, data.ClientId, data.StrategyId, data.Name, data.Version, data.CreatedAt, data.Status).Exec(); err != nil {
		return err
	}
	return nil
}
