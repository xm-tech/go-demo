package datastruct

import (
	"container/list"
	"fmt"
	"testing"
)

func setup() {
}

func TestNewLinkeList(t *testing.T) {
	list := NewLinkeList()
	list.Display()
}

func TestPush(t *testing.T) {
	list := NewLinkeList()
	list.Push(4)
	list.Push(5)
	list.Push("~~~~~6")

	list.Display()

	if list.Size() != 3 {
		t.Errorf("LinkeList.Push error, size expected 3, but got %d", list.Size())
	}
}

func TestOfficialList(t *testing.T) {
	list := list.New()
	list.PushFront("3")
	list.PushFront(3)
	list.PushFront(6)

	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}