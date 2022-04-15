package session

import (
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/squeakycheese75/go-mud/game"
	"github.com/squeakycheese75/go-mud/model"
)

type SessionHandler struct {
	dungeon      *model.Dungeon
	users        map[string]*model.User
	eventChannel chan model.SessionEvent
}

func NewSessionHandler(dungeon *model.Dungeon, eventChannel chan model.SessionEvent) *SessionHandler {
	return &SessionHandler{
		dungeon:      dungeon,
		eventChannel: eventChannel,
		users:        map[string]*model.User{},
	}
}

func generateName() string {
	return fmt.Sprintf("User %d", rand.Intn(100)+1)
}

type Game struct {
	user       *model.User
	gameEngine *game.GameHandler
	hero       *model.Hero
}

type GameCommand struct {
	valid   bool
	command string
}

func parseInstuction(message string) GameCommand {
	switch strings.ToLower(message) {
	case "c", "char", "character", "info":
		return GameCommand{valid: true, command: "character"}
	case "h", "help":
		return GameCommand{valid: true, command: "help"}
	case "i", "inventory", "inv":
		return GameCommand{valid: true, command: "inventory"}
	}

	return GameCommand{valid: false}
}

func (s *SessionHandler) Start() {
	games := make(map[string]Game)

	for sessionEvent := range s.eventChannel {
		session := sessionEvent.Session

		switch event := sessionEvent.Event.(type) {
		case *model.SessionCreatedEvent:
			// g, ok := games[session.SessionId()]
			// if ok {
			// 	log.Println("Loading gamre ")
			// 	// g.user.Session = sessionEvent.Session
			// 	break
			// }
			// Creates a new user
			user := &model.User{
				Name:    generateName(),
				Session: sessionEvent.Session}
			s.users[session.SessionId()] = user

			// Create a new Character
			hero := model.NewCharacter(user)

			// Creates a new game instamce
			game := game.NewGameHandler(user)
			games[session.SessionId()] = Game{
				user,
				game,
				hero}
			games[session.SessionId()].gameEngine.StartAdventure()
		case *model.SessionDisconnectedEvent:
			log.Println("Received SessionDisconnectedEvent")
		case *model.SessionInputEvent:
			cmd := parseInstuction(event.Msg)
			if !cmd.valid {
				games[session.SessionId()].user.Session.WriteLine("I don't understand!")
				break
			}
			switch cmd.command {
			case "character":
				games[session.SessionId()].hero.Stats()
			case "help":
				games[session.SessionId()].gameEngine.Help()
			case "inventory":
				games[session.SessionId()].hero.Inventory()
			}
		}
	}
}
