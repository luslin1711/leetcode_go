package _75_findRadius



	/*
	冬季已经来临。 你的任务是设计一个有固定加热半径的供暖器向所有房屋供暖。

	现在，给出位于一条水平线上的房屋和供暖器的位置，找到可以覆盖所有房屋的最小加热半径。

	所以，你的输入将会是房屋和供暖器的位置。你将输出供暖器的最小加热半径。

	说明:

	给出的房屋和供暖器的数目是非负数且不会超过 25000。
	给出的房屋和供暖器的位置均是非负数且不会超过10^9。
	只要房屋位于供暖器的半径内(包括在边缘上)，它就可以得到供暖。
	所有供暖器都遵循你的半径标准，加热的半径也一样。
	示例 1:

	输入: [1,2,3],[2]
	输出: 1
	解释: 仅在位置2上有一个供暖器。如果我们将加热半径设为1，那么所有房屋就都能得到供暖。
	示例 2:

	输入: [1,2,3,4],[1,4]
	输出: 1
	解释: 在位置1, 4上有两个供暖器。我们需要将加热半径设为1，这样所有房屋就都能得到供暖。

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/heaters
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


func findRadius(houses []int, heaters []int) int {
	length := len(houses)
	if length <= 1 {
		return 0
	}
	h_len := len(heaters)
	path := heaters[0] - houses[0]
	if heaters[0] - houses[0] < houses[length-1] - heaters[h_len-1] {
		path = length - heaters[h_len-1]
	}

	t_path := 0
	j := 0
	for i := 0 ; i < length; i++ {
		if houses[i] == heaters[j] {
			if j < h_len - 1 {
				j++
			}
			if t_path / 2 > path {
				path = t_path /2
			}
		}
		t_path++
	}
	return path
}

func FindRadius(houses []int, heaters []int) int {
	return findRadius(houses,heaters)
}