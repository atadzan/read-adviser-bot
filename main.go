package main

import (
	"flag"
	"log"

	tgClient "github.com/atadzan/read-adviser-bot/clients/telegram"
	"github.com/atadzan/read-adviser-bot/consumer/event-consumer"
	"github.com/atadzan/read-adviser-bot/events/telegram"
	"github.com/atadzan/read-adviser-bot/storage/files"
)

const (
	tgBotHost         = "api.telegram.org"
	sqliteStoragePath = "data/sqlite/storage.db"
	storagePath       = "storage"
	batchSize         = 100
)

func main() {
	s := files.New(storagePath)

	//if err := s.Init(context.TODO()); err != nil {
	//	log.Fatal("can't init storage: ", err)
	//}

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		s,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
