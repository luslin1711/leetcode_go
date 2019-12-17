package __ReverseInteger

func reverse(x int) int {
	const (
		MaxInt32  = 1<<31 - 1
		MinInt32  = -1 << 31
	)
	var (
		bit int
		result int
	)
	if x > -10 && x < 10 {
		return x
	}
	result = 0
	if x < 100000000 && x > -100000000 {
		for x != 0  {
			bit = x % 10
			result = result * 10 + bit
			x = x / 10
		}
	} else {
		for x != 0  {
			bit = x % 10
			if result > MaxInt32 / 10 || (result == MaxInt32 / 10 && bit > 7){
				return 0
			}
			if result < MinInt32 / 10 || (result == MinInt32 / 10 && bit < -8){
				return 0
			}
			result = result * 10 + bit
			x = x / 10
		}
	}

	return result
}

func Reverse(x int) int {
	return reverse(x)
}