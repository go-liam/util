package conv

import (
	"math/rand"
	"time"
)

// RandomNumber :
func RandomNumber(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return randNum
}
