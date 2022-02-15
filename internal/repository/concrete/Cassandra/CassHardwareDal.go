package Cassandra

import (
	"StorageWorkerService/internal/model"
	cassandra "StorageWorkerService/pkg/database/Cassandra"
	"fmt"
	"github.com/gocql/gocql"
)

type cassHardwareDal struct {
	Client *gocql.Session
	Table  string
}

func NewHardwareDal(Table string) *cassHardwareDal {
	return &cassHardwareDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassHardwareDal) Add(data *model.HardwareModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO %s(id, client_id, project_id, customer_id, device_model, device_name, device_type, graphics_device_name, graphics_device_type, graphics_device_vendor, graphics_device_version, graphics_memory_size, operating_system, processor_count, processor_type, system_memory_size, date_time, status) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId, data.DeviceModel, data.DeviceName, data.GraphicsDeviceType, data.GraphicsDeviceVendor, data.GraphicsDeviceVendor, data.GraphicsDeviceVersion, data.GraphicsMemorySize, data.OperatingSystem, data.ProcessorCount, data.ProcessorType, data.SystemMemorySize, data.DateTime, data.Status).Exec(); err != nil {
		return err
	}
	return nil
}
