package Cassandra

import (
	"StorageWorkerService/internal/model"
	cassandra "StorageWorkerService/pkg/database/Cassandra"
	"fmt"
	"github.com/gocql/gocql"
)

type cassScreenSwipeDal struct {
	Client *gocql.Session
	Table  string
}

func NewScreenSwipeDal(Table string) *cassScreenSwipeDal {
	return &cassScreenSwipeDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassScreenSwipeDal) Add(data *model.ScreenSwipeModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO %s(id, client_id, project_id, customer_id, start_loc_x, start_loc_y, finish_loc_x, finish_loc_y, level_name, level_index, swipe_direction, created_at, status) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId, data.StartLocX, data.StartLocY, data.FinishLocX, data.FinishLocY, data.LevelName, data.LevelIndex, data.SwipeDirection, data.CreatedAt, data.Status).Exec(); err != nil {
		return err
	}
	return nil
}
