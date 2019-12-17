package _005_largestSumAfterKNegations

/*
	给定一个整数数组 A，我们只能用以下方法修改该数组：我们选择某个个索引 i 并将 A[i] 替换为 -A[i]，然后总共重复这个过程 K 次。（我们可以多次选择同一个索引 i。）
	以这种方式修改数组后，返回数组可能的最大和。
	示例 1：

	输入：A = [4,2,3], K = 1
	输出：5
	解释：选择索引 (1,) ，然后 A 变为 [4,-2,3]。
	示例 2：

	输入：A = [3,-1,0,2], K = 3
	输出：6
	解释：选择索引 (1, 2, 2) ，然后 A 变为 [3,1,0,2]。
	示例 3：

	输入：A = [2,-3,-1,5,-4], K = 2
	输出：13
	解释：选择索引 (1, 4) ，然后 A 变为 [2,3,-1,5,4]。
	提示：
	1 <= A.length <= 10000
	1 <= K <= 10000
	-100 <= A[i] <= 100
	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/maximize-sum-of-array-after-k-negations
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func largestSumAfterKNegations(A []int, K int) int {
	var (
		minAbs,ans,i int
	)
	QuickSort(A,0, len(A)-1)
	minAbs = 101
	for i = 0; i < len(A); i++ {
		if A[i] < 0{
			if K > 0 {
				ans += -A[i]
				K--
			} else {
				ans += A[i]
			}
			if -A[i] < minAbs{
				minAbs = -A[i]
			}

		} else {
			ans += A[i]
			if A[i] < minAbs{
				minAbs = A[i]
			}
		}
	}
	if K % 2 == 1 {
		ans -= 2*minAbs
	}
	return ans
}



func LargestSumAfterKNegations(A []int, K int) int {
	return largestSumAfterKNegations(A,K)
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