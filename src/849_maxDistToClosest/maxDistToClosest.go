package _49_maxDistToClosest

/*
	在一排座位（ seats）中，1 代表有人坐在座位上，0 代表座位上是空的。

	至少有一个空座位，且至少有一人坐在座位上。

	亚历克斯希望坐在一个能够使他与离他最近的人之间的距离达到最大化的座位上。

	返回他到离他最近的人的最大距离。

	示例 1：

	输入：[1,0,0,0,1,0,1]
	输出：2
	解释：
	如果亚历克斯坐在第二个空位（seats[2]）上，他到离他最近的人的距离为 2 。
	如果亚历克斯坐在其它任何一个空位上，他到离他最近的人的距离为 1 。
	因此，他到离他最近的人的最大距离是 2 。
	示例 2：

	输入：[1,0,0,0]
	输出：3
	解释：
	如果亚历克斯坐在最后一个座位上，他离最近的人有 3 个座位远。
	这是可能的最大距离，所以答案是 3 。
	提示：

	1 <= seats.length <= 20000
	seats 中只含有 0 和 1，至少有一个 0，且至少有一个 1。

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/maximize-distance-to-closest-person
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func maxDistToClosest(seats []int) int {
	left_max, middle_max,right_max := 0,0,0
	length := len(seats)
	p := 0
	q := length-1
	for p = 0; p < length; p++ {
		if seats[p] == 1 {
			break
		}
		left_max++
	}
	for q = length-1; q >= 0; q-- {
		if seats[q] == 1 {
			break
		}
		right_max++
	}
	tmp := 0
	for ; p <= q; p++ {
		if seats[p] == 1 {
			middle_max = max2(tmp, middle_max)
			tmp = 0
		} else {
			tmp++
		}
	}

	return max3(left_max, (middle_max + 1) / 2, right_max)
}
func max2(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func max3(a,b,c int) int {
	if a >= b && a >= c {
		return  a
	} else if b >= a && b >= c {
		return b
	} else {
		return c
	}
}