package Cassandra

import (
	"StorageWorkerService/internal/model"
	cassandra "StorageWorkerService/pkg/database/Cassandra"
	"fmt"
	"github.com/gocql/gocql"
)

type cassScreenClickDal struct {
	Client *gocql.Session
	Table  string
}

func NewScreenClickDal(Table string) *cassScreenClickDal {
	return &cassScreenClickDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassScreenClickDal) Add(data *model.ScreenClickModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO %s(id, client_id, project_id, customer_id, start_loc_x, start_loc_y, finish_loc_x, finish_loc_y, level_name, level_index, tab_count, finger_id, created_at, status) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId, data.StartLocX, data.StartLocY, data.FinishLocX, data.FinishLocY, data.LevelName, data.LevelIndex, data.TabCount, data.FingerId, data.CreatedAt, data.Status).Exec(); err != nil {
		return err
	}
	return nil
}
