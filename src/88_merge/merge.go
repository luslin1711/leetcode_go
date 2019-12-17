package _8_merge

/*
	给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。

	说明:

	初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
	你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
	示例:

	输入:
	nums1 = [1,2,3,0,0,0], m = 3
	nums2 = [2,5,6],       n = 3

	输出: [1,2,2,3,5,6]

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/merge-sorted-array
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func merge(nums1 []int, m int, nums2 []int, n int)  {
	var (
		i,j,p int
	)
	i = m - 1
	j = n - 1
	for p = m + n - 1; p>=0 && i >=0 && j>=0; p-- {
		if nums1[i] > nums2[j] {
			nums1[p] = nums1[i]
			i--
		} else {
			nums1[p] = nums2[j]
			j--
		}
	}
	for ; i>=0; i-- {
		nums1[p] = nums1[i]
		p--
	}
	for ; j>=0; j-- {
		nums1[p] = nums2[j]
		p--
	}

	nums1 = nums1[:m+n]

}

func Merge(nums1 []int, m int, nums2 []int, n int)  {
	merge(nums1, m, nums2, n)
}