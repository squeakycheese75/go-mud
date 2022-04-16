package game

import (
	"fmt"

	"github.com/squeakycheese75/go-mud/data"
	"github.com/squeakycheese75/go-mud/model"
)

type GameHandler struct {
	story    *data.DungeonData
	user     *model.User
	location uint
	stage    data.Stage
}

func NewGameHandler(user *model.User) *GameHandler {
	return &GameHandler{
		story:    data.LoadData(),
		user:     user,
		location: 0,
	}
}

func (h *GameHandler) LoadStage(page int) {
	for _, stage := range h.story.Stages {
		if stage.Page == page {
			h.stage = stage
			h.user.Session.WriteLine(h.stage.Narrative)
			h.ShowChoices(stage.Options)
			return
		}
	}
	h.WriteLine("Error loading the next stage")
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

func (h *GameHandler) Show() {
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
	h.LoadStage(page)
}
