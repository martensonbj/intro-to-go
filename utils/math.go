package utils

func Add(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	return total
}
