package arraygenerate

import (
	"math/rand"
)

func ArrayGenerate(value int) []int {
	size := rand.Intn(10)
	numbers := make([]int, size)
	for i := 0; i < size; i++ {
		numbers[i] = rand.Intn(10000) + 1
	}
	return numbers
}
