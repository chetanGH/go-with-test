package arrays

func SumArr(array []int) int {
	var sum int
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}
