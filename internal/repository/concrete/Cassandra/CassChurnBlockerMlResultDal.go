package Cassandra

import (
	"StorageWorkerService/internal/model"
	cassandra "StorageWorkerService/pkg/database/Cassandra"
	"fmt"
	"github.com/gocql/gocql"
)

type cassChurnBlockerMlResultDal struct {
	Client *gocql.Session
	Table  string
}

func NewCassChurnBlockerMlResultDal(Table string) *cassChurnBlockerMlResultDal {
	return &cassChurnBlockerMlResultDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassChurnBlockerMlResultDal) Add(data *model.ChurnBlockerMlResultModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO %s(id, client_id, project_id, customer_id, model_type, model_result, date_time, status) VALUES(?,?,?,?,?,?,?,?)", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId, data.ModelType, data.ModelResult, data.DateTime, data.Status).Exec(); err != nil {
		return err
	}
	return nil
}
