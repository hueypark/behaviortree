package behaviortree

// BehaviorTree is a tree of nodes for represent AI.
type BehaviorTree struct {
	rootNode *Node
}

// New creates a new BehaviorTree.
func New() *BehaviorTree {
	bt := &BehaviorTree{}

	return bt
}
