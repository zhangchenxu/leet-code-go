package main

import "fmt"

type CQueue struct {
	stk1, stk2 []int
}

func Constructor() CQueue {
	return CQueue{[]int{}, []int{}}
}

func (this *CQueue) appendTail(val int) {
	this.stk1 = append(this.stk1, val)
}

func (this *CQueue) deleteHead() int {
	if len(this.stk2) == 0 {
		if len(this.stk1) != 0 {
			for _, v := range this.stk1 {
				this.stk2 = append(this.stk2, v)
				this.stk1 = this.stk1[:len(this.stk1)-1]
			}
		}
	}
	if len(this.stk2) == 0 {
		return -1
	}
	val := this.stk2[len(this.stk2)-1]
	this.stk2 = this.stk2[:len(this.stk2)-1]
	return val
}

func main() {
	queue := Constructor()
	queue.appendTail(1)
	queue.appendTail(2)
	queue.appendTail(3)

	fmt.Println(queue.deleteHead())
}
