package dice

import (
	"math/rand"
)

func RollDice() int {
	return (rand.Intn(5) + 1)
}
