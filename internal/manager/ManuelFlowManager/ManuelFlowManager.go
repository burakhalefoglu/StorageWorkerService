package ManuelFlowManager

import (
	"context"
	"log"
	"time"

	mongodb "StorageWorkerService/internal/dataAccess/mongodb"
	jsonParser "StorageWorkerService/internal/helper/jsonParser"

	"github.com/segmentio/kafka-go"
)

type ManuelFlowModel struct{

	ClientId string
	ProjectId string
	CustomerId string
	DifficultyLevel int
	DateTime time.Time
}


func InsertManuelFlowModel(reader *kafka.Reader, m kafka.Message) {

	manuelFlowModel := ManuelFlowModel{}
		jsonParser.DecodeJson(m.Value, &manuelFlowModel)
		log.Println(manuelFlowModel)
	
	_, err:= mongodb.AddCollection(m.Topic, manuelFlowModel)
		if(err == nil) {
				if err := reader.CommitMessages(context.Background(), m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}
}