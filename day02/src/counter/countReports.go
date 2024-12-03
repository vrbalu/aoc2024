package counter

func CountSafeReports(allReports [][]int) int {
	var reportSafe int
	for _, report := range allReports {
		var safeIncreasing, safeDecreasing int
		for j := 0; j < len(report)-1; j++ {
			if isIncreasingWithMaximumDistance3(report[j], report[j+1]) {
				safeIncreasing++
				continue
			}
			if isDecreasingithMaximumDistance3(report[j], report[j+1]) {
				safeDecreasing++
			}
		}
		// -1 because of number of comaprison between numbers is one smaller as the total amount
		if safeIncreasing == len(report)-1 || safeDecreasing == len(report)-1 {
			reportSafe++
		}
	}
	return reportSafe
}

func isIncreasingWithMaximumDistance3(num1, num2 int) bool {
	if num1 > num2 {
		if num1-num2 <= 3 {
			return true
		}
	}
	return false
}
func isDecreasingithMaximumDistance3(num1, num2 int) bool {
	if num1 < num2 {
		if (num2 - num1) <= 3 {
			return true
		}
	}
	return false
}
