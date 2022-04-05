package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"tcms-web-bridge/internal/connections/kafka"
	"tcms-web-bridge/internal/dry"
	tcms2 "tcms-web-bridge/internal/tcms"
	"tcms-web-bridge/internal/telegramClient"
	"tcms-web-bridge/internal/webserver"
)

func main() {
	log := logrus.New()
	// Load values from .env into the system
	err := godotenv.Load()
	if err != nil {
		log.Error(err)
	}

	telegram, err := telegramClient.NewTelegram()
	dry.HandleErrorPanic(err)

	addConsumer := make(chan chan []uint8)
	quitKafka := make(chan bool)
	kafkaError := make(chan error)

	go kafka.CreateKafkaSubscription(addConsumer, kafkaError, quitKafka)

	tcms, err := tcms2.GetTcms()
	if err != nil {
		panic(err)
	}
	go webserver.StartWebServer(telegram, tcms, addConsumer)

	select {}
}
