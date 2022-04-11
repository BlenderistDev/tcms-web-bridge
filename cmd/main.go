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

	telegram, err := getTelegramClient(&config)
	if err != nil {
		panic(err)
	}

	addConsumer := make(chan chan []uint8)
	quitKafka := make(chan bool)
	kafkaError := make(chan error)

	go kafka.CreateKafkaSubscription(config, addConsumer, kafkaError, quitKafka)

	tcms, err := getTcmsClient(&config)
	if err != nil {
		panic(err)
	}

	go webserver.StartWebServer(config, telegram, tcms, addConsumer)

	select {}
}

func getTcmsClient(config *config2.Config) (tcms2.Tcms, error) {
	tcmsConn, err := grpc.Dial(config.TcmsHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	tcms := tcms2.GetTcms(tcmsConn)
	return tcms, nil
}

func getTelegramClient(config *config2.Config) (telegramClient.TelegramClient, error) {
	conn, err := grpc.Dial(config.TelegramBridgeHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	tg := telegramClient.GetTelegram(conn)
	return tg, nil
}
