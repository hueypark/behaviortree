package behaviortree

import "testing"

func TestBehaviorTreeEqual(t *testing.T) {
	bt1 := New()
	bt2 := New()
	if !bt1.Equal(bt2) {
		t.Fatalf("Not implemented")
	}
}
