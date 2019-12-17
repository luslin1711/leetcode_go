package _19_containsNearbyDuplicate


/*
	给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的绝对值最大为 k。

	示例 1:

	输入: nums = [1,2,3,1], k = 3
	输出: true
	示例 2:

	输入: nums = [1,0,1,1], k = 1
	输出: true
	示例 3:

	输入: nums = [1,2,3,1,2,3], k = 2
	输出: false
 */

func containsNearbyDuplicate(nums []int, k int) bool {
	bitMap := make(map[int]int)
	for i:=0; i< len(nums);i++ {
		value , exist := bitMap[nums[i]]
		if ! exist {
			bitMap[nums[i]] = i
		} else {
			if i - value <= k {
				return true
			} else {
				bitMap[nums[i]] = i
			}
		}
	}
	return false
}

func ContainsNearbyDuplicate(nums []int, k int) bool {
	return containsNearbyDuplicate(nums,k)
}