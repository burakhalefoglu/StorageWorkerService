package model

import "time"

type HardwareModel struct{

	ClientId string
	ProjectId string
	CustomerId string
	DeviceModel string
	DeviceName string
	DeviceType int
	GraphicsDeviceName string
	GraphicsDeviceType int
	GraphicsDeviceVendor string
	GraphicsDeviceVersion string
	GraphicsMemorySize int
	OperatingSystem string
	ProcessorCount int
	ProcessorType string
	SystemMemorySize int
	DateTime time.Time
}
