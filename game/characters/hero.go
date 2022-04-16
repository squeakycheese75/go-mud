package characters

import (
	"fmt"

	"github.com/squeakycheese75/go-mud/game/dice"
	"github.com/squeakycheese75/go-mud/model"
)

func NewHero(u *model.User) *Hero {
	// Roll die
	skill := (dice.RollDice() + 6)
	stamina := (dice.RollDice() + dice.RollDice() + 12)
	luck := (dice.RollDice() + 6)

	hero := &Hero{
		user: u,
		base: PlayerStats{
			Skill:   skill,
			Stamina: stamina,
			Luck:    luck,
		},
		Stats: PlayerStats{
			Skill:   skill,
			Stamina: stamina,
			Luck:    luck,
		},
		inventory:  Inventory{},
		experience: Experience{monsterVanquished: 0},
	}
	return hero
}

type Inventory struct {
	Items      []interface{}
	Gold       int
	Jewels     []interface{}
	Potions    []interface{}
	Provisions []interface{}
}

type Experience struct {
	monsterVanquished int
}

type Hero struct {
	user       *model.User
	base       PlayerStats
	Stats      PlayerStats
	inventory  Inventory
	experience Experience
}

type PlayerStats struct {
	Skill   int
	Stamina int
	Luck    int
}

type Character interface {
	Alive()
}

func (h *Hero) Alive() bool {
	return h.Stats.Stamina > 0
}

func (h *Hero) WriteLine(msg string) {
	h.user.Session.WriteLine(msg)
}

func (h *Hero) ShowStats() {
	h.WriteLine("*** Character Info ***")
	h.WriteLine(fmt.Sprintf("Name: %v", h.user.Name))
	h.WriteLine("Stats:")
	h.WriteLine(fmt.Sprintf(" - Skill: %v of %v", h.Stats.Skill, h.base.Skill))
	h.WriteLine(fmt.Sprintf(" - Luck: %v of %v", h.Stats.Luck, h.base.Luck))
	h.WriteLine(fmt.Sprintf(" - Stamina: %v of %v", h.Stats.Stamina, h.base.Stamina))

	h.ShowInventory()
	h.ShowInfo()
}

func (h *Hero) ShowInventory() {
	h.WriteLine("Inventory:")
	h.WriteLine(fmt.Sprintf("Gold: %v", h.inventory.Gold))
	// h.WriteLine(fmt.Sprintf("Jewels: %v", h.inventory.Jewels...))
	// h.WriteLine(fmt.Sprintf("Items: %v", h.inventory.Items...))
	// h.WriteLine(fmt.Sprintf("Potions: %v", h.inventory.Potions...))
}

func (h *Hero) ShowInfo() {
	h.WriteLine("Experience:")
	h.WriteLine(fmt.Sprintf(" - Monsters Vanquished:  %v", h.experience.monsterVanquished))
}

func (h *Hero) TakeHit() {
	h.Stats.Stamina = (h.Stats.Stamina - 2)
}

func (h *Hero) FoeDefeated() {
	h.experience.monsterVanquished = h.experience.monsterVanquished + 1
}
