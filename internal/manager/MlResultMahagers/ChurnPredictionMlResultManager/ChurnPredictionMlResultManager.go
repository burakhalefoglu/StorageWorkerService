package ChurnPredictionMlResultManager

import (
	"context"
	"log"
	"time"

	mongodb "StorageWorkerService/internal/dataAccess/mongodb"
	jsonParser "StorageWorkerService/internal/helper/jsonParser"

	"github.com/segmentio/kafka-go"
)

type ChurnPredictionMlResultModel struct{

	ClientId string
	ProjectId string
	CustomerId string
	ModelType string
	ModelResult float32
	DateTime time.Time
}


func InsertChurnPredictionMlResultModel(reader *kafka.Reader, m kafka.Message) {

	churnPredictionMlResultModel := ChurnPredictionMlResultModel{}
		jsonParser.DecodeJson(m.Value, &churnPredictionMlResultModel)
		log.Println(churnPredictionMlResultModel)
	
	_, err:= mongodb.AddCollection(m.Topic, churnPredictionMlResultModel)
		if(err == nil) {
				if err := reader.CommitMessages(context.Background(), m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}
}