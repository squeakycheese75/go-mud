package session

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/squeakycheese75/go-mud/game"
	"github.com/squeakycheese75/go-mud/model"
)

type SessionHandler struct {
	dungeon      *model.Dungeon
	users        map[string]*model.User
	eventChannel chan model.SessionEvent
	games        map[string]*model.Adventure
}

func NewSessionHandler(dungeon *model.Dungeon, eventChannel chan model.SessionEvent) *SessionHandler {
	return &SessionHandler{
		dungeon:      dungeon,
		eventChannel: eventChannel,
		users:        map[string]*model.User{},
		games:        map[string]*model.Adventure{},
	}
}

func generateName() string {
	return fmt.Sprintf("User %d", rand.Intn(100)+1)
}

func (s *SessionHandler) Start() {
	for sessionEvent := range s.eventChannel {
		session := sessionEvent.Session

		switch event := sessionEvent.Event.(type) {

		case *model.SessionCreatedEvent:
			// Creates a new user
			user := &model.User{
				Name:    generateName(),
				Session: sessionEvent.Session}
			s.users[session.SessionId()] = user

			// Create a new Character
			model.NewCharacter(user)

			// Creates a new game instamce
			game := game.NewGameHandler(user)
			game.StartAdventure()

		case *model.SessionDisconnectedEvent:
			log.Println("Received SessionDisconnectedEvent")
		case *model.SessionInputEvent:
			log.Printf("Received SessionInputEvent: \"%s\" from User: %s\r\n", event.Msg, event.Msg)
		}
	}
}
