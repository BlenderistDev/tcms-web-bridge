package kafka

import (
	"testing"

	"tcms-web-bridge/internal/dry"
)

func TestGetKafkaHost(t *testing.T) {
	dry.TestEnvString(t, "KAFKA_HOST", getKafkaHost)
}

func TestGetKafkaTopic(t *testing.T) {
	dry.TestEnvString(t, "KAFKA_TOPIC", getKafkaTopic)
}

func TestGetKafkaGroupId(t *testing.T) {
	dry.TestEnvString(t, "KAFKA_GROUP_ID", getKafkaGroupId)
}
