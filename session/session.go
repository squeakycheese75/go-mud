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

func ParseUserInstuction(message string, options []data.Option) (interface{}, error) {
	msg := strings.ToLower(message)
	// Check if it's a game command
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
	// Check if it's a stage option.
	for _, v := range options {
		if msg == v.Key {
			return GameChoice{stage: v.Next}, nil
		}
	}
	return nil, errors.New("unknown user instruction")
}

func (s *SessionHandler) Start() {
	// games := make(map[string]Game)
	games := make(map[string]*game.GameHandler)

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
			s.users[session.SessionId()] = sessionEvent.User

			// Creates a new game instamce
			session.WriteLine(fmt.Sprintln("Begin your Quest...."))
			game := game.NewGameHandler(user, characters.NewHero(user))
			game.LoadStage(1)

			games[session.SessionId()] = game
		case *model.SessionDisconnectedEvent:
			log.Println("Received SessionDisconnectedEvent")
		case *model.SessionInputEvent:
			game := games[session.SessionId()]
			ins, err := ParseUserInstuction(event.Msg, games[session.SessionId()].GetOptions())
			if err != nil {
				game.WriteLine("I don't understand!  Do you want (h)elp?")
				break
			}
			switch ins := ins.(type) {
			case GameCommand:
				switch ins.command {
				case "character":
					game.Hero.ShowStats()
				case "help":
					game.Help()
				case "inventory":
					game.Hero.ShowInventory()
				case "location":
					game.ShowStage()
				}
			case GameChoice:
				game.NextStage(ins.stage)
			}
		}
		// Check
		session.Write(">> ")
	}
}
