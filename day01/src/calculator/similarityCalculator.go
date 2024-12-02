package calculator

func CalculateSimiliarity(leftArray, rightArray []int) int {
	var similiarity int
	similiarityMap := make(map[int]int)
	for _, number := range leftArray {
		similiarityMap[number] = 0
	}
	for _, number := range rightArray {
		if _, ok := similiarityMap[number]; ok {
			similiarityMap[number]++
		}
	}
	for _, number := range leftArray {
		amount := similiarityMap[number]
		similiarity += number * amount
	}
	return similiarity
}
