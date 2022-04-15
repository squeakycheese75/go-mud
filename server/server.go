package server

import (
	"log"
	"net"

	"github.com/squeakycheese75/go-mud/model"
)

// func generateName() string {
// 	return fmt.Sprintf("User %d", rand.Intn(100)+1)
// }

// func HandleConnection(conn net.Conn, inputChannel chan model.ClientInput) error {
// 	log.Println("Connected!")
// 	buf := make([]byte, 4096)

// 	session := &model.Session{Conn: conn}
// 	user := &model.User{
// 		Name:    generateName(),
// 		Session: session}

// 	inputChannel <- model.ClientInput{
// 		User:    user,
// 		Session: session,
// 		Event:   &model.UserJoinedEvent{}}

// 	for {
// 		n, err := conn.Read(buf)
// 		if err != nil {
// 			return err
// 		}
// 		if n == 0 {
// 			log.Println("Zero bytes, disconnected")
// 			break
// 		}
// 		msg := buf[0 : n-2]
// 		if !game.CheckInput(string(msg)) {
// 			session.WriteLine("Invalid Input")
// 			continue
// 		}
// 		inputChannel <- model.ClientInput{
// 			User:    user,
// 			Session: session,
// 			Event: &model.MessageEvent{
// 				Msg: string(msg),
// 			}}

// 		// response := fmt.Sprintf("You said: \"%s\"\r\n", msg)
// 		// n, err = conn.Write([]byte(response))
// 		// if err != nil {
// 		// 	return err
// 		// }
// 		if n == 0 {
// 			log.Println("Zero bytes, disconnected")
// 			break
// 		}
// 	}
// 	return nil
// }

func SessionHandleConnection(conn net.Conn, sessionEventChannel chan model.SessionEvent) error {
	log.Println("Session established!")
	buf := make([]byte, 4096)
	session := &model.Session{Conn: conn}
	// user := &model.User{
	// 	Name:    generateName(),
	// 	Session: session}

	sessionEventChannel <- model.SessionEvent{
		Session: session,
		Event:   &model.SessionCreatedEvent{}}
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
		// go func() {
		// 	if err := HandleConnection(conn, clientInputChannel); err != nil {
		// 		log.Println("Error handling connection", err)
		// 		return
		// 	}
		// }()
		go func() {
			if err := SessionHandleConnection(conn, sessionEventChannel); err != nil {
				log.Println("Error handling connection", err)
				return
			}
		}()
	}
}
