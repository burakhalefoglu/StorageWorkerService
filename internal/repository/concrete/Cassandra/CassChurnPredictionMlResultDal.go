package Cassandra

import (
	"StorageWorkerService/internal/model"
	cassandra "StorageWorkerService/pkg/database/Cassandra"
	"fmt"
	"github.com/gocql/gocql"
)

type cassChurnPredictionMlResultDal struct {
	Client *gocql.Session
	Table  string
}

func NewCassChurnPredictionMlResultDal(Table string) *cassChurnPredictionMlResultDal {
	return &cassChurnPredictionMlResultDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassChurnPredictionMlResultDal) Add(data *model.ChurnPredictionMlResultModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO %s(id, client_id, project_id, customer_id, model_type, model_result, date_time, status) VALUES(?,?,?,?,?,?,?,?)", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId, data.ModelType, data.ModelResult, data.DateTime, data.Status).Exec(); err != nil {
		return err
	}
	return nil
}
