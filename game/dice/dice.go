package dice

import (
	"math/rand"
	"time"
)

func RollDice() int {
	rand.Seed(time.Now().UnixNano())
	return (rand.Intn(5) + 1)
}
