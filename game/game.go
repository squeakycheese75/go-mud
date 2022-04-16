package game

import (
	"fmt"

	"github.com/squeakycheese75/go-mud/data"
	"github.com/squeakycheese75/go-mud/game/characters"
	"github.com/squeakycheese75/go-mud/game/combat"
	"github.com/squeakycheese75/go-mud/model"
)

type GameHandler struct {
	story    *data.DungeonData
	user     *model.User
	location uint
	stage    data.Stage
	Hero     *characters.Hero
}

func NewGameHandler(user *model.User, hero *characters.Hero) *GameHandler {
	return &GameHandler{
		story:    data.LoadData(),
		user:     user,
		location: 0,
		Hero:     hero,
	}
}

func (h *GameHandler) LoadStage(stageNo int) error {
	for _, stage := range h.story.Stages {
		if stage.Page == stageNo {
			h.stage = stage
			h.ShowStage()
			return nil
		}
	}
	return fmt.Errorf("stage %v not found", stageNo)
}

func (h *GameHandler) ShowChoices(options []data.Option) {
	h.WriteLine("")
	h.WriteLine("Options: ")
	for _, v := range options {
		h.WriteLine(fmt.Sprintf(" - %s", v.Choice))
	}
	h.WriteLine("")
	h.WriteLine(fmt.Sprintln("Choose:"))
}

func (h *GameHandler) GetOptions() []data.Option {
	return h.stage.Options
}

func (h *GameHandler) ShowStage() {
	h.WriteLine("")
	h.WriteLine(h.stage.Narrative)
	h.WriteLine("")
	// Show actions
	switch h.stage.Action {
	case "choose":
		h.ShowChoices(h.stage.Options)
	}
}

func (h *GameHandler) Help() {
	h.WriteLine("Type:")
	h.WriteLine("(l) - show current location")
	h.WriteLine("(c) - show characteter info")
	h.WriteLine("(i) - show characteter inventory")
	h.WriteLine("")
}

func (h *GameHandler) WriteLine(msg string) {
	h.user.Session.WriteLine(msg)
}

func (h *GameHandler) Write(msg string) {
	h.user.Session.Write(msg)
}

func (h *GameHandler) NextStage(page int) {
	h.location = uint(page)
	err := h.LoadStage(page)
	if err != nil {
		h.user.Session.WriteLine(err.Error())
	}
	if h.stage.Action == "fight" {
		h.user.Session.WriteLine("We are about to fight!!!!!!!!!!!")
		for _, c := range h.stage.Characters {
			h.user.Session.WriteLine(fmt.Sprintf("It's a %s", c.Name))
			c := combat.NewBattle(h.Hero, characters.NewMonster(c.Name, c.Stats.Stamina, c.Stats.Skill), h.user)
			c.Fight()
		}
	}
}
