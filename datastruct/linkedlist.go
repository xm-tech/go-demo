package datastruct

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

var lock sync.Mutex

type LinkedList struct {
	head *Node
	size int
}

func NewLinkeList() *LinkedList {
	return &LinkedList{head: nil, size: 0}
}

// A linkedlist Node
type Node struct {
	val  int
	next *Node
}

// add one element into the linkedlist head
func (self *LinkedList) Push(val int) {
	lock.Lock()
	defer lock.Unlock()

	node := &Node{val: val, next: nil}
	if self.head == nil {
		self.head = node
		self.size++
		return
	}

	old_head := self.head
	self.head = node
	self.head.next = old_head
	self.size++
}

// delete the last element, and return the first element
func (self *LinkedList) pop() *Node {
	return nil
}

// iterate display the all elements of the in the linkedlist
func (self *LinkedList) Display() {
	if self.head == nil {
		fmt.Println("Empty LinkeList")
		return
	}

	// loop print every element
	var output strings.Builder
	node := self.head
	for {
		output.WriteString(strconv.Itoa(node.val))
		output.WriteString(".")
		if node.next == nil {
			break
		} else {
			node = node.next
		}
	}

	fmt.Println(output.String())
}

func (self *LinkedList) Size() int {
	return self.size
}
