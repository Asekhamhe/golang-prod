package dinner

import (
	"math/rand"
	"time"
)

// Choose is function that returns a dinner
func Choose() string {
	dinners := []string{
		"tacos",
		"pizza",
		"ramen",
	}

	rand.Seed(time.Now().Unix())

	return dinners[rand.Intn(len(dinners))]
}
