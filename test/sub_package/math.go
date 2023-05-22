package sub

func Sum(args ...int) int {
	sum := 0
	for _, n := range args {
		sum += n
	}
	return sum
}
