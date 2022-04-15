package model

import (
	"fmt"

	"github.com/squeakycheese75/go-mud/utils"
)

func NewCharacter(user *User) *Hero {
	user.Session.WriteLine("*** New Character Created ***")
	stamina := utils.RollDice()
	hero := &Hero{
		user:       user,
		skill:      utils.RollDice(),
		stamina:    stamina,
		luck:       utils.RollDice(),
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
	user       *User
	skill      int
	luck       int
	stamina    int
	life       int
	inventory  []Inventory
	experience Experience
}

type Character interface {
	Alive()
}

func (hero Hero) Alive() bool {
	return hero.life > 0
}

func (hero Hero) Stats() {
	hero.user.Session.WriteLine("*** Character Info ***")
	hero.user.Session.WriteLine(fmt.Sprintf("Name: %v", hero.user.Name))
	hero.user.Session.WriteLine("Stats:")
	hero.user.Session.WriteLine(fmt.Sprintf(" - Skill: %v", hero.skill))
	hero.user.Session.WriteLine(fmt.Sprintf(" - Luck: %v", hero.luck))
	hero.user.Session.WriteLine(fmt.Sprintf(" - Stamina: %v of %v", hero.life, hero.stamina))
	hero.Inventory()
	hero.Info()
}

func (hero Hero) Inventory() {
	hero.user.Session.WriteLine("Inventory:")
	for _, v := range hero.inventory {
		hero.user.Session.WriteLine(fmt.Sprintf(" - %v x %v", v.quantity, v.Name))
	}
}

func (hero Hero) Info() {
	hero.user.Session.WriteLine("Experience:")
	hero.user.Session.WriteLine(fmt.Sprintf(" - Monsters Vanquished:  %v", hero.experience.monsterVanquished))
}
