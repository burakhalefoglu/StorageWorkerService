package ChurnBlokerMlResultManager

import (
	"context"
	"log"
	"time"

	mongodb "StorageWorkerService/internal/dataAccess/mongodb"
	jsonParser "StorageWorkerService/internal/helper/jsonParser"

	"github.com/segmentio/kafka-go"
)

type ChurnBlokerMlResultModel struct{

	ClientId string
	ProjectId string
	CustomerId string
	ModelType string
	ModelResult float32
	DateTime time.Time
}


func InsertChurnBlokerMlResultModel(reader *kafka.Reader, m kafka.Message) {

	churnBlokerMlResultModel := ChurnBlokerMlResultModel{}
		jsonParser.DecodeJson(m.Value, &churnBlokerMlResultModel)
		log.Println(churnBlokerMlResultModel)
	
	_, err:= mongodb.AddCollection(m.Topic, churnBlokerMlResultModel)
		if(err == nil) {
				if err := reader.CommitMessages(context.Background(), m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}
}