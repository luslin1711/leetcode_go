package util

import "fmt"

type ArrayQueue struct {
	q        []interface{}
	capacity int
	head     int
	tail     int
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{make([]interface{}, n), n, 0, 0}
}

func (this *ArrayQueue) Push(v interface{}) bool {
	if this.tail == this.capacity {
		return false
	}
	this.q[this.tail] = v
	this.tail++
	return true
}

func (this *ArrayQueue) Pop() interface{} {
	if this.head == this.tail {
		return nil
	}
	v := this.q[this.head]
	this.head++
	return v
}

func (this *ArrayQueue) Print() {
	if this.head == this.tail {
		fmt.Println("empty queue")
	}
	result := "head"
	for i := this.head; i < this.tail; i++ {
		result += fmt.Sprintf("<-%+v", this.q[i])
	}
	result += "<-tail"
	fmt.Println(result)
}

