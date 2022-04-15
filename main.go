package main

import (
	"log"

	"github.com/squeakycheese75/go-mud/model"
	"github.com/squeakycheese75/go-mud/server"
	"github.com/squeakycheese75/go-mud/session"
)

func main() {
	dungeon := &model.Dungeon{}

	eventChannel := make(chan model.SessionEvent)

	session := session.NewSessionHandler(dungeon, eventChannel)
	go session.Start()

	err := server.StartServer(eventChannel)
	if err != nil {
		log.Fatal("Error starting the server")
	}
}
