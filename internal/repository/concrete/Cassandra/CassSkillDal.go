package Cassandra

import (
	"StorageWorkerService/internal/model"
	cassandra "StorageWorkerService/pkg/database/Cassandra"
	"fmt"
	"github.com/gocql/gocql"
)

type cassSkillDal struct {
	Client *gocql.Session
	Table  string
}

func NewSkillDal(Table string) *cassSkillDal {
	return &cassSkillDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassSkillDal) Add(data *model.SkillModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO %s(id, inventory_id, skill_type, count) VALUES(?,?,?,?)", m.Table),
		data.Id, data.InventoryId, data.SkillType, data.Count).Exec(); err != nil {
		return err
	}
	return nil
}
