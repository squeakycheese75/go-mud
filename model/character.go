package model

import (
	"fmt"

	"github.com/squeakycheese75/go-mud/utils"
)

func NewCharacter(user *User) *Hero {
	user.Session.WriteLine("*** New Character Created ***")
	hero := &Hero{
		user:    user,
		skill:   utils.RollDice(),
		stamina: utils.RollDice(),
		luck:    utils.RollDice(),
	}
	hero.Stats()
	return hero
}

type Hero struct {
	user    *User
	skill   int
	luck    int
	stamina int
}

// type Monster struct {
// 	name    string
// 	skill   int
// 	luck    int
// 	stamina int
// }

type Character interface {
	Alive()
}

func (hero Hero) Alive() bool {
	return hero.stamina > 0
}

// func (monster Monster) Alive() bool {
// 	return monster.stamina > 0
// }

func (hero Hero) Health() {
	hero.user.Session.WriteLine(fmt.Sprintf("Stamina: %v", hero.stamina))
}

func (hero Hero) Stats() {
	hero.user.Session.WriteLine(fmt.Sprintf("Name: %v", hero.user.Name))
	hero.user.Session.WriteLine(fmt.Sprintf("Luck: %v", hero.luck))
	hero.user.Session.WriteLine(fmt.Sprintf("Skill: %v\r\n", hero.skill))
}
