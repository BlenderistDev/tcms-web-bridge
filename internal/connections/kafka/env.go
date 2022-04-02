package kafka

import (
	"tcms-web-bridge/internal/dry"
)

func getKafkaHost() (string, error) {
	return dry.GetEnvStr("KAFKA_HOST")
}

func getKafkaTopic() (string, error) {
	return dry.GetEnvStr("KAFKA_TOPIC")
}

func getKafkaGroupId() (string, error) {
	return dry.GetEnvStr("KAFKA_GROUP_ID")
}
