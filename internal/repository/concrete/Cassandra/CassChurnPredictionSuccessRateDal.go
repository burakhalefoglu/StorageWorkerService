package Cassandra

import (
	"StorageWorkerService/internal/model"
	cassandra "StorageWorkerService/pkg/database/Cassandra"
	"fmt"

	"github.com/gocql/gocql"
)

type cassChurnPredictionSuccessRateDal struct {
	Client *gocql.Session
	Table  string
}

func NewCassChurnPredictionSuccessRateDal(Table string) *cassChurnPredictionSuccessRateDal {
	return &cassChurnPredictionSuccessRateDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassChurnPredictionSuccessRateDal) Add(data *model.ChurnPredictionSuccessRateModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO %s(id, project_id, value, created_at, status) VALUES(?,?,?,?,?)", m.Table),
		data.Id, data.ProjectId, data.Value, data.CreatedAt, data.Status).Exec(); err != nil {
		return err
	}
	return nil
}
