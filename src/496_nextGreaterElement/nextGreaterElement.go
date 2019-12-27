package _96_nextGreaterElement

import (
	"fmt"
	"math"
)

/*
	给定两个没有重复元素的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。找到 nums1 中每个元素在 nums2 中的下一个比其大的值。

	nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出-1。

	示例 1:

	输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
	输出: [-1,3,-1]
	解释:
		对于num1中的数字4，你无法在第二个数组中找到下一个更大的数字，因此输出 -1。
		对于num1中的数字1，第二个数组中数字1右边的下一个较大数字是 3。
		对于num1中的数字2，第二个数组中没有下一个更大的数字，因此输出 -1。
	示例 2:

	输入: nums1 = [2,4], nums2 = [1,2,3,4].
	输出: [3,-1]
	解释:
	    对于num1中的数字2，第二个数组中的下一个较大数字是3。
		对于num1中的数字4，第二个数组中没有下一个更大的数字，因此输出 -1。
	注意:

	nums1和nums2中所有元素是唯一的。
	nums1和nums2 的数组大小都不超过1000。

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/next-greater-element-i
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type ArrayStack struct {
	data []int
	top  int // 栈顶指针
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{make([]int, 0), -1}
}

func (this *ArrayStack) IsEmpty() bool {
	if this.top < 0 {
		return true
	}
	return false
}

func (this *ArrayStack) Push(v int) {
	if this.top < 0 {
		this.top = 0
	} else {
		this.top++
	}
	if this.top > len(this.data)-1 {
		this.data = append(this.data, v)
	} else {
		this.data[this.top] = v
	}
}

func (this *ArrayStack) Pop() int {
	if this.IsEmpty() {
		return math.MinInt32
	}
	v := this.data[this.top]
	this.top--
	return v
}

func (this *ArrayStack) Top() int {
	if this.IsEmpty() {
		return math.MinInt32
	}
	return this.data[this.top]
}

func (this *ArrayStack) Flush() {
	this.top = -1
}

func (this *ArrayStack) Print() {
	if this.IsEmpty() {
		fmt.Println("empty statck")
	} else {
		for i := this.top; i >= 0; i-- {
			fmt.Println(this.data[i])
		}
	}
}
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if len(nums2) <= 1 {
		for i := 0; i < len(nums1); i++ {
			nums1[i] = -1
		}
		return nums1
	}
	n2_map := make(map[int]int)
	stack := NewArrayStack()
	stack.Push(nums2[0])
	for i := 1; i < len(nums2); i++ {
		for nums2[i] > stack.Top() && !stack.IsEmpty() {
			n2_map[stack.Pop()] = nums2[i]

		}
		stack.Push(nums2[i])
	}
	for i := 0; i < len(nums1); i++ {
		value, ok := n2_map[nums1[i]]
		if !ok {
			nums1[i] = -1
		} else {
			nums1[i] = value
		}
	}
	return nums1
}


func NextGreaterElement(nums1 []int, nums2 []int) []int {
	return nextGreaterElement(nums1,nums2)
}