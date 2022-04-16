package combat

import (
	"fmt"

	"github.com/squeakycheese75/go-mud/game/characters"
	"github.com/squeakycheese75/go-mud/game/dice"
	"github.com/squeakycheese75/go-mud/model"
)

type Battle struct {
	hero    *characters.Hero
	monster *characters.Monster
	user    *model.User
}

func NewBattle(hero *characters.Hero, monster *characters.Monster, user *model.User) *Battle {
	return &Battle{
		hero:    hero,
		monster: monster,
		user:    user,
	}
}

func (c *Battle) Fight() bool {
	var round int = 1
	for c.hero.Alive() && c.monster.Alive() {
		c.user.Session.WriteLine(fmt.Sprintf("Round %v", round))
		// Monster roll
		monsterAttack := dice.RollDice() + dice.RollDice() + c.monster.Skill
		c.user.Session.WriteLine(fmt.Sprintf("Monster Attack Score: %v ", monsterAttack))
		// Hero Roll
		heroAttack := dice.RollDice() + dice.RollDice() + c.hero.Stats.Skill
		c.user.Session.WriteLine(fmt.Sprintf("Hero Attack Score: %v ", heroAttack))

		if monsterAttack > heroAttack {
			c.user.Session.WriteLine(fmt.Sprintf("The %v deals a mighty blow", c.monster.Name))
			c.hero.TakeHit()
		} else if heroAttack > monsterAttack {
			c.user.Session.WriteLine(fmt.Sprintf("You deal a mighty blow to the %v", c.monster.Name))
			c.monster.TakeHit()
		} else {
			c.user.Session.WriteLine(fmt.Sprintf("Swing and a miss"))
		}

		round = round + 1
		c.user.Session.WriteLine(fmt.Sprintf("Your stamina %v", c.hero.Stats.Stamina))
		c.user.Session.WriteLine(fmt.Sprintf("%v stamina %v", c.monster.Name, c.monster.Stamina))
	}
	if c.hero.Alive() {
		// c.hero.
		c.user.Session.WriteLine("You stand victorious!!")
		return true
	}
	c.user.Session.WriteLine("You're dead!")
	return false
}
