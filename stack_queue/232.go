package stackqueue

type MyQueue struct {
	data []int
}

func Constructor() MyQueue {
	return MyQueue{data: make([]int, 0)}
}

func (this *MyQueue) Push(x int) {
	this.data = append(this.data, x)
}

func (this *MyQueue) Pop() int {
	res := this.data[0]
	this.data = this.data[1:]
	return res
}

func (this *MyQueue) Peek() int {
	return this.data[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.data) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
