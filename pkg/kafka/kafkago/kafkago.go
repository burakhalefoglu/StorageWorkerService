package kafkago

import (
	"StorageWorkerService/pkg/helper"
	"context"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

type kafkaGo struct {
}

func KafkaGoConstructor() *kafkaGo {
	return &kafkaGo{}
}

func (k *kafkaGo) Produce(key *[]byte, value *[]byte, topic string) (err error) {
	writer, _ := writerConfigure([]string{helper.ResolvePath("KAFKA_HOST", "KAFKA_PORT")}, uuid.New().String(), topic)
	message := kafka.Message{
		Key:   *key,
		Value: *value,
		Time:  time.Now(),
	}
	err = writer.WriteMessages(context.Background(), message)
	log.Print("kafkaGo", "Producer", topic, key)
	return err
}

func (k *kafkaGo) Consume(topic string, groupId string, wg *sync.WaitGroup, callback func(data *[]byte) (bool, string)) {
	reader, _ := readerConfigure([]string{helper.ResolvePath("KAFKA_HOST", "KAFKA_PORT")}, groupId, topic)
	log.Println("Consumer Started: ", topic, groupId)
	defer func(reader *kafka.Reader) {
		err := reader.Close()
		if err != nil {
			log.Fatal("kafkaGo", "Consume", "failed to reader.Close() messages:"+err.Error())
		}
	}(reader)
	log.Print("kafkaGo", "Consume", reader.Stats().ClientID)
	for {
		m, err := reader.FetchMessage(context.Background())
		if err != nil {
			log.Fatal("kafkaGo", "Consume", "error while receiving message: "+err.Error())
			continue
		}
		log.Println("kafkaGo", "Consume", topic, groupId)
		isSuccess, _ := callback(&m.Value)
		if isSuccess {
			if err := reader.CommitMessages(context.Background(), m); err != nil {
				log.Fatal("kafkaGo", "Consume", "failed to commit messages:"+err.Error())
			}
		}
	}
	wg.Done()
}

func writerConfigure(kafkaBrokerUrls []string, clientId string, topic string) (w *kafka.Writer, err error) {
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: clientId,
	}

	config := kafka.WriterConfig{
		Brokers:          kafkaBrokerUrls,
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}
	w = kafka.NewWriter(config)
	return w, nil
}

func readerConfigure(kafkaBrokerUrls []string, groupID string, topic string) (r *kafka.Reader, err error) {
	config := kafka.ReaderConfig{
		Brokers:         kafkaBrokerUrls,
		GroupID:         groupID,
		Topic:           topic,
		MinBytes:        10e3,            // 10KB
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka_go.
		ReadLagInterval: -1,
	}

	reader := kafka.NewReader(config)
	return reader, nil
}
