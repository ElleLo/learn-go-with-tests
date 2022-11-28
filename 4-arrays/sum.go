package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// ... variadic functions, any number of trailing arguments
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	// lengthOfNumbers := len(numbersToSum)

	// make() is used to create a slice, with a starting capacity of the len of the numbersToSum
	// sums := make([]int, lengthOfNumbers)

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// Calculating the Tails of each slice
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			// slice[low:high], in this case: take from 1 to the end
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}

	}

	return sums

}
