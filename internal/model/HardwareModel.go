package model

import "time"

type HardwareModel struct {
	Id                    int64
	ClientId              int64
	ProjectId             int64
	CustomerId            int64
	DeviceModel           string
	DeviceName            string
	DeviceType            int
	GraphicsDeviceName    string
	GraphicsDeviceType    int
	GraphicsDeviceVendor  string
	GraphicsDeviceVersion string
	GraphicsMemorySize    int
	OperatingSystem       string
	ProcessorCount        int
	ProcessorType         string
	SystemMemorySize      int
	DateTime              time.Time
	Status                bool
}
