package calculator

import (
	"math"
	"sort"
)

func CalculateDistance(leftArray, rightArray []int) int {
	var distance int
	sortedLeft, sortedRight := sortArrays(leftArray, rightArray)

	for i := 0; i < len(sortedLeft); i++ {
		distance += int(math.Abs(float64(sortedLeft[i]) - float64(sortedRight[i])))
	}
	return distance
}

func sortArrays(left, right []int) ([]int, []int) {
	sort.Ints(left)
	sort.Ints(right)
	return left, right

}
