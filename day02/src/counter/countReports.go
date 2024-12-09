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
		} else if (len(report)-1-safeIncreasing) <= 2 || (len(report)-1-safeDecreasing) <= 2 {
			reportSafe += handleSpecial(report)
		}
	}
	return reportSafe
}

func handleSpecial(report []int) int {
	for i := 0; i < len(report); i++ {
		var safeIncreasing, safeDecreasing int
		reportCopy := report
		copy(reportCopy[i:], reportCopy[i+1:])
		reportCopy[len(reportCopy)-1] = 0 // or the zero value of T
		reportCopy = reportCopy[:len(reportCopy)-1]
		for j := 0; j < len(reportCopy)-1; j++ {
			if isIncreasingWithMaximumDistance3(reportCopy[j], reportCopy[j+1]) {
				safeIncreasing++
				continue
			}
			if isDecreasingithMaximumDistance3(reportCopy[j], reportCopy[j+1]) {
				safeDecreasing++
			}
		}
		if safeIncreasing == len(reportCopy)-1 || safeDecreasing == len(reportCopy)-1 {
			return 1
		}
	}
	return 0
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
