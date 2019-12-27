package _24_pivotIndex

func pivotIndex(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return 0
	}
	p := 0
	q := length -1
	sump := 0
	sumq := 0
	for q > p {
		if sump  > sumq {
			sumq += nums[q]
			q--
		} else if sump < sumq {
			sump += nums[p]
			p++
		} else {
			sumq += nums[q]
			q--
		}
	}
	if sump == sumq {
		return p
	}
	return -1
}