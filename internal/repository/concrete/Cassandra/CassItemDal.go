package Cassandra

import (
	"StorageWorkerService/internal/model"
	cassandra "StorageWorkerService/pkg/database/Cassandra"
	"fmt"
	"github.com/gocql/gocql"
)

type cassItemDal struct {
	Client *gocql.Session
	Table  string
}

func NewItemDal(Table string) *cassItemDal {
	return &cassItemDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassItemDal) Add(data *model.ItemModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO %s(id, inventory_id, item_type, count) VALUES(?,?,?,?)", m.Table),
		data.Id, data.InventoryId, data.ItemType, data.Count).Exec(); err != nil {
		return err
	}
	return nil
}
