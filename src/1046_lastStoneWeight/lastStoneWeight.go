package _046_lastStoneWeight

import (
	"fmt"
)

/*
	有一堆石头，每块石头的重量都是正整数。
	每一回合，从中选出两块最重的石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：
	如果 x == y，那么两块石头都会被完全粉碎；
	如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
	最后，最多只会剩下一块石头。返回此石头的重量。如果没有石头剩下，就返回 0。
	提示：
	1 <= stones.length <= 30
	1 <= stones[i] <= 1000
	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/last-stone-weight
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func lastStoneWeight(stones []int) int {
	var (
		length, i, max1, max2 int
		heap *Heap
	)

	length = len(stones)
	heap = NewHeap(length)
	for i=0; i< length; i++ {
		heap.Push(stones[i])
	}
	fmt.Println(heap.a)
	for heap.Len() > 1 {
		max1 = heap.Pop()
		max2 = heap.Pop()
		fmt.Println(max2,max1)
		if max1 - max2 != 0 {
			heap.Push(max1 - max2)
		}
	}
	return heap.Pop()
}

func LastStoneWeight(stones []int) int {
	return lastStoneWeight(stones)
}

type Heap struct {
	a     []int
	n     int
	count int
}

//init heap
func NewHeap(capacity int) *Heap {
	heap := &Heap{}
	heap.n = capacity
	heap.a = make([]int, capacity+1)
	heap.count = 0
	return heap
}

//top-max heap -> heapify from down to up
func (heap *Heap) Push(data int) {
	//defensive
	if heap.count == heap.n {
		return
	}

	heap.count++
	heap.a[heap.count] = data

	//compare with parent node
	i := heap.count
	parent := i / 2
	for parent > 0 && heap.a[parent] < heap.a[i] {
		swap(heap.a, parent, i)
		i = parent
		parent = i / 2
	}
}
func (heap *Heap) Len() int {
	return heap.count
}

//heapfify from up to down
func (heap *Heap) Pop() int {

	//defensive
	if heap.count == 0 {
		return 0
	}
	max := heap.a[1]
	//swap max and last
	swap(heap.a, 1, heap.count)
	heap.count--

	//heapify from up to down
	heapifyUpToDown(heap.a, heap.count)
	return max
}

//heapify
func heapifyUpToDown(a []int, count int) {

	for i := 1; i <= count/2; {

		maxIndex := i
		if a[i] < a[i*2] {
			maxIndex = i * 2
		}

		if i*2+1 <= count && a[maxIndex] < a[i*2+1] {
			maxIndex = i*2 + 1
		}

		if maxIndex == i {
			break
		}

		swap(a, i, maxIndex)
		i = maxIndex
	}

}

//swap two elements
func swap(a []int, i int, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}