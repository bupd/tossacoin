package randomint

import (
	"math/rand/v2"
)

func RandomInt(num uint8) int {
	counts := make([]int, num)

	for i := 0; i < 69; i++ { // Number of times to run randSelect
		randomChoice := rand.IntN(int(num)) // Generates a random integer between 0 and choices-1
		counts[randomChoice]++
	}

	maxCount := 0
	maxChoice := 0
	for i, count := range counts {
		if count > maxCount {
			maxCount = count
			maxChoice = i
		}
	}

	return maxChoice
}
