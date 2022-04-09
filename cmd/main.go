package main

import (
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	config2 "tcms-web-bridge/internal/config"
	"tcms-web-bridge/internal/connections/kafka"
	tcms2 "tcms-web-bridge/internal/tcms"
	"tcms-web-bridge/internal/telegramClient"
	"tcms-web-bridge/internal/webserver"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	config, err := config2.LoadConfig(wd)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	telegram, err := telegramClient.NewTelegram(config)
	if err != nil {
		panic(err)
	}

	addConsumer := make(chan chan []uint8)
	quitKafka := make(chan bool)
	kafkaError := make(chan error)

	go kafka.CreateKafkaSubscription(config, addConsumer, kafkaError, quitKafka)

	tcmsConn, err := grpc.Dial(config.TcmsHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	tcms, err := tcms2.GetTcms(tcmsConn)
	if err != nil {
		panic(err)
	}
	go webserver.StartWebServer(config, telegram, tcms, addConsumer)

	select {}
}
