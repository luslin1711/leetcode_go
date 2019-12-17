package _029_twoCitySchedCost


/*
	公司计划面试 2N 人。第 i 人飞往 A 市的费用为 costs[i][0]，飞往 B 市的费用为 costs[i][1]。

	返回将每个人都飞到某座城市的最低费用，要求每个城市都有 N 人抵达。
	示例：

	输入：[[10,20],[30,200],[400,50],[30,20]]
	输出：110
	解释：
	第一个人去 A 市，费用为 10。
	第二个人去 A 市，费用为 30。
	第三个人去 B 市，费用为 50。
	第四个人去 B 市，费用为 20。

	最低总费用为 10 + 30 + 50 + 20 = 110，每个城市都有一半的人在面试。
	 

	提示：

	1 <= costs.length <= 100
	costs.length 为偶数
	1 <= costs[i][0], costs[i][1] <= 1000

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/two-city-scheduling
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


func twoCitySchedCost(costs [][]int) int {
	var (
		i int
		ans int
	)
	QuickSort(costs,0, len(costs)-1)
	for i=0; i < len(costs) / 2; i++ {
		ans += costs[i][0]
	}
	for ;i < len(costs); i++ {
		ans += costs[i][1]
	}
	return ans
}

func TwoCitySchedCost(costs [][]int) int {
	return twoCitySchedCost(costs)
}

func point(arr [][]int, left, right int) int {
	var (
		start, end int
	)
	start = left
	end = right
	tmpInt := arr[left][0] - arr[left][1]
	for start < end {
		for ;start < end && arr[end][0] - arr[end][1] >= tmpInt ; end--{}
		for ;start < end && arr[start][0] - arr[start][1] <= tmpInt; start++{}

		arr[start], arr[end] = arr[end], arr[start]
	}
	arr[start] , arr[left] = arr[left], arr[start]
	return start
}

func QuickSort(arr [][]int, left, right int)  {
	if left < right {
		i := point(arr, left, right)
		QuickSort(arr, left, i -1)
		QuickSort(arr, i+1, right)
	}
}