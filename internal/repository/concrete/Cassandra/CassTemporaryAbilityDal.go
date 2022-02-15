package Cassandra

import (
	"StorageWorkerService/internal/model"
	cassandra "StorageWorkerService/pkg/database/Cassandra"
	"fmt"
	"github.com/gocql/gocql"
)

type cassTemporaryAbilityDal struct {
	Client *gocql.Session
	Table  string
}

func NewTemporaryAbilityDal(Table string) *cassTemporaryAbilityDal {
	return &cassTemporaryAbilityDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassTemporaryAbilityDal) Add(data *model.TemporaryAbilityModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO %s(id, inventory_id, ability_type, count) VALUES(?,?,?,?)", m.Table),
		data.Id, data.InventoryId, data.AbilityType, data.Count).Exec(); err != nil {
		return err
	}
	return nil
}
