package datastruct

import "testing"

func setup() {
}

func TestNewNode(t *testing.T) {
	head := NewNode(3)
	if head.val != 3 {
		t.Errorf("NewNode error, val expedted 3 bug got %d", head.val)
	}
}

func TestPush(t *testing.T) {
	head := NewNode(3)
	head.push(4)
	if head.next == nil {
		t.Errorf("Node.push error, the next not set")
	}

	if head.next.val != 4 {
		t.Errorf("Node.push error, expedted 4 but got %d", head.next.val)
	}
}

func TestPop(t *testing.T) {
	head := NewNode(3)
	head.push(4)
	poped := head.pop()
	if poped == nil {
		t.Errorf("Node.pop error, no element")
	}
	if poped.val != 4 {
		t.Errorf("Node.pop error, expected 4 but got %d", poped.val)
	}
}
