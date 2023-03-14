package datastruct

// A linkedlist Node
type Node struct {
	val  int
	next *Node
}

func NewNode(val int) *Node {
	return &Node{val: val, next: nil}
}

// add one element into the linkedlist head
func (self *Node) push(val int) {
	self.next = &Node{val: val}
}

// delete the last element, and return the first element
func (self *Node) pop() *Node {

	return &Node{val: self.val}
}

// iterate display the all elements of the in the linkedlist
func (self *Node) display() {

}
