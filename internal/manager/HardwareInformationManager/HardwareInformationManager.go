package HardwareInformationManager

import (
	"context"
	"log"
	"time"

	mongodb "StorageWorkerService/internal/dataAccess/mongodb"
	jsonParser "StorageWorkerService/internal/helper/jsonParser"

	"github.com/segmentio/kafka-go"
)

type HardwareInformationModel struct{

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


func InsertHardwareInformationModel(reader *kafka.Reader, m kafka.Message) {

	hardwareInformationModel := HardwareInformationModel{}
		jsonParser.DecodeJson(m.Value, &hardwareInformationModel)
		log.Println(hardwareInformationModel)
	
	_, err:= mongodb.AddCollection(m.Topic, hardwareInformationModel)
		if(err == nil) {
				if err := reader.CommitMessages(context.Background(), m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}
}