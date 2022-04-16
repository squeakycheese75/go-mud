package session

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/google/uuid"
	"github.com/squeakycheese75/go-mud/data"
	"github.com/squeakycheese75/go-mud/game"
	"github.com/squeakycheese75/go-mud/game/characters"
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
	hero       *characters.Hero
}

type GameCommand struct {
	valid   bool
	command string
}

type GameChoice struct {
	stage int
}

func NameGenerator() string {
	return fmt.Sprintf("user %v", rand.Intn(100))
}

func parseInstuction(message string, options []data.Option) (interface{}, error) {
	msg := strings.ToLower(message)
	switch msg {
	case "c", "char", "character", "info":
		return GameCommand{command: "character", valid: true}, nil
	case "h", "help":
		return GameCommand{command: "help", valid: true}, nil
	case "i", "inventory", "inv":
		return GameCommand{command: "inventory", valid: true}, nil
	case "l", "location", "loc":
		return GameCommand{command: "location", valid: true}, nil
	}

	for _, v := range options {
		if msg == v.Key {
			return GameChoice{stage: v.Next}, nil
		}

	}
	return nil, errors.New("unknown message")
}

func (s *SessionHandler) Start() {
	games := make(map[string]Game)

	for sessionEvent := range s.eventChannel {
		session := sessionEvent.Session

		switch event := sessionEvent.Event.(type) {
		case *model.SessionCreatedEvent:
			// Create a new user
			user := &model.User{
				ID:      uuid.New(),
				Name:    NameGenerator(),
				Session: session}

			log.Printf("\"%v\" joined", user.ID)
			session.WriteLine(fmt.Sprintf("\"%v\" joined", user.ID))
			s.users[session.SessionId()] = sessionEvent.User
			// log.Printf("SessionId %v", session.SessionId())

			// Create a new Character
			hero := characters.NewHero(user)

			// Creates a new game instamce
			session.WriteLine(fmt.Sprintln("Begin your Quest...."))

			game := game.NewGameHandler(user)
			game.LoadStage(1)
			games[session.SessionId()] = Game{user, game, hero}
		case *model.SessionDisconnectedEvent:
			log.Println("Received SessionDisconnectedEvent")
		case *model.SessionInputEvent:
			log.Printf("Received SessionInputEvent: %v", event.Msg)
			ins, err := parseInstuction(event.Msg, games[session.SessionId()].gameEngine.GetOptions())
			if err != nil {
				games[session.SessionId()].user.Session.WriteLine("I don't understand!")
				break
			}
			switch ins := ins.(type) {
			case GameCommand:
				log.Println("It's a GameCommand")
				switch ins.command {
				case "character":
					games[session.SessionId()].hero.Stats()
				case "help":
					games[session.SessionId()].gameEngine.Help()
				case "inventory":
					games[session.SessionId()].hero.Inventory()
				case "location":
					games[session.SessionId()].gameEngine.Show()
				}
			case GameChoice:
				log.Println("It's a GameChoice")
				games[session.SessionId()].gameEngine.NextStage(ins.stage)

			}
		}
		session.Write(">> ")
	}
}
