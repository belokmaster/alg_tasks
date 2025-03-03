package main

/*
type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head *Node
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}

	current := this.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{val: val, next: this.head}
	this.head = node
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{val: val}
	if this.head == nil {
		this.head = node
	} else {
		current := this.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	curr := this.head
	for i := 0; i < index-1; i++ {
		curr = curr.next
	}

	node := &Node{val: val, next: curr.next}
	curr.next = node
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}

	if index == 0 {
		this.head = this.head.next
	} else {
		curr := this.head
		for i := 0; i < index-1; i++ {
			curr = curr.next
		}
		curr.next = curr.next.next
	}
	this.size--
}
/*
/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
*/
