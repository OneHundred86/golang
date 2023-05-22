package sub_package

func Sum(args ...int) int {
	sum := 0
	for _, n := range args {
		sum += n
	}
	return sum
}
