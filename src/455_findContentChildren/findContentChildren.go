package _55_findContentChildren


/*
	假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。对每个孩子 i ，都有一个胃口值 gi ，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j ，都有一个尺寸 sj 。如果 sj >= gi ，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。

	注意：

	你可以假设胃口值为正。
	一个小朋友最多只能拥有一块饼干。

	示例 1:

	输入: [1,2,3], [1,1]

	输出: 1

	解释:
	你有三个孩子和两块小饼干，3个孩子的胃口值分别是：1,2,3。
	虽然你有两块小饼干，由于他们的尺寸都是1，你只能让胃口值是1的孩子满足。
	所以你应该输出1。
	示例 2:

	输入: [1,2], [1,2,3]

	输出: 2

	解释:
	你有两个孩子和三块小饼干，2个孩子的胃口值分别是1,2。
	你拥有的饼干数量和尺寸都足以让所有孩子满足。
	所以你应该输出2.

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/assign-cookies
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


// 贪心
func findContentChildren(g []int, s []int) int {
	var (
		g_len, s_len int
		i,j int
	)
	g_len = len(g)
	s_len = len(s)
	if g_len == 0 || s_len == 0 {
		return 0
	}
	QuickSort(g,0,g_len-1)
	QuickSort(s,0,s_len-1)
	i = 0
	j = 0
	for i < g_len && j < s_len {
		if s[j] >= g[i] {
			i++
		}
		j++
	}
	return i
}

func FindContentChildren1(g []int, s []int) int {
	return findContentChildren(g,s)
}

func point(arr []int, left, right int) int {
	var (
		start, end int
	)
	start = left
	end = right
	tmpInt := arr[left]
	for start < end {
		for ;start < end && arr[end] >= tmpInt ; end--{}
		for ;start < end && arr[start] <= tmpInt; start++{}

		arr[start], arr[end] = arr[end], arr[start]
	}
	arr[start] , arr[left] = arr[left], arr[start]
	return start
}

func QuickSort(arr []int, left, right int)  {
	if left < right {
		i := point(arr, left, right)
		QuickSort(arr, left, i -1)
		QuickSort(arr, i+1, right)
	}
}