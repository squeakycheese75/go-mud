package characters

import (
	"fmt"

	"github.com/squeakycheese75/go-mud/game/dice"
	"github.com/squeakycheese75/go-mud/model"
)

func NewHero(u *model.User) *Hero {
	u.Session.WriteLine("*** New Character Created ***")
	stamina := dice.RollDice()

	hero := &Hero{
		user:       u,
		skill:      dice.RollDice(),
		stamina:    stamina,
		luck:       dice.RollDice(),
		life:       stamina,
		inventory:  []Inventory{},
		experience: Experience{monsterVanquished: 0},
	}
	hero.inventory = append(hero.inventory, Inventory{Name: "Healing Potion", quantity: 1})

	hero.Stats()
	return hero
}

type Inventory struct {
	Name     string
	quantity int
}

type Experience struct {
	monsterVanquished int
}

type Hero struct {
	user       *model.User
	skill      int
	luck       int
	stamina    int
	life       int
	inventory  []Inventory
	experience Experience
}

type Character interface {
	Alive()
	WriteLine(msg string)
}

func (h *Hero) Alive() bool {
	return h.life > 0
}

func (h *Hero) WriteLine(msg string) {
	h.user.Session.WriteLine(msg)
}

func (h *Hero) Stats() {
	h.WriteLine("*** Character Info ***")
	h.WriteLine(fmt.Sprintf("Name: %v", h.user.Name))
	h.WriteLine("Stats:")
	h.WriteLine(fmt.Sprintf(" - Skill: %v", h.skill))
	h.WriteLine(fmt.Sprintf(" - Luck: %v", h.luck))
	h.WriteLine(fmt.Sprintf(" - Stamina: %v of %v", h.life, h.stamina))

	h.Inventory()
	h.Info()
}

func (h *Hero) Inventory() {
	h.WriteLine("Inventory:")
	for _, v := range h.inventory {
		h.WriteLine(fmt.Sprintf(" - %v x %v", v.quantity, v.Name))
	}
}

func (h *Hero) Info() {
	h.WriteLine("Experience:")
	h.WriteLine(fmt.Sprintf(" - Monsters Vanquished:  %v", h.experience.monsterVanquished))
}
