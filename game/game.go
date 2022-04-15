package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/squeakycheese75/go-mud/model"
)

type GameHandler struct {
	story *model.Story
	user  *model.User
}

func NewGameHandler(user *model.User) *GameHandler {
	return &GameHandler{
		story: readFromJson(),
		user:  user,
	}
}

func (s *GameHandler) getStoryStage(page int) *model.Stage {
	var resval = s.story.Stages[page]
	return &resval
}

func (s *GameHandler) StartAdventure() {
	s.user.Session.WriteLine(fmt.Sprintln("Entering the Dungeon"))
	s.Display(0)
}

func (s *GameHandler) Display(page int) {
	stage := s.getStoryStage(page)
	s.user.Session.WriteLine(fmt.Sprint(stage.Narrative))
	s.user.Session.WriteLine(fmt.Sprintln("Choose:"))
	for _, v := range stage.Options {
		s.user.Session.WriteLine(fmt.Sprintf("Options: %s", v.Choice))
	}
}

func readFromJson() *model.Story {
	// Open our jsonFile
	jsonFile, err := os.Open("./game/dungeon.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var story model.Story

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &story)
	return &story
}

func (s *GameHandler) Help() {
	s.user.Session.WriteLine("Boohoo, you need help!")
}
