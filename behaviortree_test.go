package behaviortree

import "testing"

func TestBehaviorTreeEqual(t *testing.T) {
	btStr := `Sequence
`
	bt1 := New(btStr)
	bt2 := New(btStr)
	if !bt1.Equal(bt2) {
		t.Fatalf("bt1 and bt2 is not equal")
	}
}

func TestBehaviorTreeNew(t *testing.T) {
	btStr := `Sequence
`

	bt := New(btStr)

	if bt.rootNode == nil {
		t.Fatalf("bt.rootNode is nil")
	}

	if !bt.rootNode.Equal(&Sequence{}) {
		t.Fatalf("bt.rootNode is not equal to Sequence")
	}
}
