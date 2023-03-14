package datastruct

import "testing"

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
	list.Push(6)

	list.Display()

	if list.Size() != 3 {
		t.Errorf("LinkeList.Push error, size expected 3, but got %d", list.Size())
	}
}
