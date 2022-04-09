package kafka

import (
	"context"
	"strings"
	"time"

	kafka2 "github.com/segmentio/kafka-go"
	"tcms-web-bridge/internal/config"
)

// CreateKafkaSubscription create subscription for kafka topic
func CreateKafkaSubscription(config config.Config, addConsumer chan chan []uint8, errChan chan error, quit chan bool) {
	var consumers []chan []uint8

	brokers := strings.Split(config.KafkaHost, ",")
	reader := kafka2.NewReader(kafka2.ReaderConfig{
		Brokers:           brokers,
		GroupID:           config.KafkaGroupId,
		Topic:             config.KafkaTopic,
		MaxBytes:          10e6, // 10MB
		MaxWait:           time.Millisecond * 10,
		HeartbeatInterval: 1,
		ReadBackoffMax:    time.Millisecond * 100,
	})

	ctx := context.Background()
	for {
		select {
		case <-quit:
			return
		case newConsumer := <-addConsumer:
			consumers = append(consumers, newConsumer)
		default:
			m, err := reader.ReadMessage(ctx)
			if err != nil {
				errChan <- err
			}
			for _, ch := range consumers {
				ch <- m.Value
			}
		}
	}
}
