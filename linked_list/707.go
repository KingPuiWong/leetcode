package linked_list

type SingleNode struct {
	Val  int
	Next *SingleNode
}

type MyLinkedList struct {
	dummyHead *SingleNode
	size      int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		dummyHead: &SingleNode{
			Val:  0,
			Next: nil,
		},
		size: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if this == nil || this.size == 0 || index >= this.size || index < 0 {
		return -1
	}

	cur := this.dummyHead.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &SingleNode{
		Val:  val,
		Next: nil,
	}

	newNode.Next = this.dummyHead.Next
	this.dummyHead.Next = newNode
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	dummy := this.dummyHead
	for dummy.Next != nil {
		dummy = dummy.Next
	}

	dummy.Next = &SingleNode{
		Val:  val,
		Next: nil,
	}
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}

	if index < 0 {
		index = 0
	}

	pre := this.dummyHead
	// 找到要插入位置的前一个节点
	for i := 0; i < index; i++ {
		pre = pre.Next
	}

	next := pre.Next
	newNode := SingleNode{
		Val:  val,
		Next: next,
	}

	pre.Next = &newNode
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size || index < 0 {
		return
	}

	//找到要删除位置的前一个节点
	pre := this.dummyHead
	for i := 0; i < index; i++ {
		pre = pre.Next
	}
	//next := pre.Next.Next
	//pre.Next = next
	if pre.Next != nil {
		pre.Next = pre.Next.Next
	}
	this.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
