package server

import (
	"fmt"
	"log"
	"net"

	"github.com/squeakycheese75/go-mud/model"
)

var nextSessionId = 1

func generateSessionId() string {
	var sid = nextSessionId
	nextSessionId++
	return fmt.Sprintf("%d", sid)
}

func SessionHandleConnection(conn net.Conn, sessionEventChannel chan model.SessionEvent) error {
	log.Println("New connection established!")
	buf := make([]byte, 4096)
	session := &model.Session{Conn: conn, Id: generateSessionId()}

	sessionEventChannel <- model.SessionEvent{
		Session: session,
		Event:   &model.SessionCreatedEvent{},
	}
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return err
		}
		if n == 0 {
			log.Println("Zero bytes, disconnected")
			sessionEventChannel <- model.SessionEvent{
				Session: session,
				Event:   &model.SessionDisconnectedEvent{}}
			break
		}
		msg := buf[0 : n-2]
		// log.Printf("Message Received: \"%v\" from user: \"%v\"", string(msg), user.Name)

		sessionEventChannel <- model.SessionEvent{
			Session: session,
			Event: &model.SessionInputEvent{
				Msg: string(msg),
			}}
	}
	return nil
}

func StartServer(sessionEventChannel chan model.SessionEvent) error {
	log.Println("Starting Server")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting connection", err)
			continue
		}

		go func() {
			if err := SessionHandleConnection(conn, sessionEventChannel); err != nil {
				log.Println("Error handling connection", err)
				return
			}
		}()
	}
}
