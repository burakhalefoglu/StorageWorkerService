package InventoryDataManager

import (
	"context"
	"log"
	"time"

	mongodb "StorageWorkerService/internal/dataAccess/mongodb"
	jsonParser "StorageWorkerService/internal/helper/jsonParser"

	"github.com/segmentio/kafka-go"
)

type InventoryDataModel struct{

	ClientId string
	ProjectId string
	CustomerId string
	MinorMine float32
	ModerateMine float32
	PreciousMine float32
	Items []Item
	Skills []Item
	TemporaryAbilitys []Item
	CreatedAt time.Time
}

type Item struct {
	ItemType string
	Count int
}

type Skill struct {
	SkillType string
	Count int
}

type TemporaryAbility struct{
	AbilityType string
	Count int
}


func InsertInventoryDataModel(reader *kafka.Reader, m kafka.Message) {

	inventoryDataModel := InventoryDataModel{}
		jsonParser.DecodeJson(m.Value, &inventoryDataModel)
		log.Println(inventoryDataModel)
	
	_, err:= mongodb.AddCollection(m.Topic, inventoryDataModel)
		if(err == nil) {
				if err := reader.CommitMessages(context.Background(), m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}
}