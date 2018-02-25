package ds

//LLNode is a node for linked list
type LLNode struct {
	Value int
	Prev  *LLNode
	Next  *LLNode
}

//LinkedList is am implementation of a Doubly linked list
type LinkedList struct {
	Head *LLNode
	Tail *LLNode
}

//NewLinkedList creates and returns a linkedlist object
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

//Find finds the first node with value v in the linked list if it exists, returns nil otherwise
func (l *LinkedList) Find(v int) *LLNode {
	if l.Head == nil {
		return nil
	}

	node := l.Head
	for node.Next != nil {
		if node.Value == v {
			return node
		}
		node = node.Next
	}

	return nil
}

//Insert inserts a new node with value v at the end of the linked list
func (l *LinkedList) Insert(v int) {
	node := &LLNode{
		Value: v,
		Prev:  l.Tail,
	}

	l.Tail.Next = node
	l.Tail = node

}

//Remove removes the first node with value v if it exists, does nothing otherwise
func (l *LinkedList) Remove(v int) {
	node := l.Find(v)
	if node == nil {
		return
	}

	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
}
