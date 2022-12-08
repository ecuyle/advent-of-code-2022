package utils

func Sum(slice []int) int {
	sum := 0

	for _, value := range slice {
		sum += value
	}

	return sum
}
