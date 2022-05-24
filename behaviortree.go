package behaviortree

import (
	"strings"
)

// BehaviorTree is a tree of nodes for represent AI.
type BehaviorTree struct {
	rootNode Node
}

// New creates a new BehaviorTree.
func New(str string) *BehaviorTree {
	bt := &BehaviorTree{}

	bt.init(str)

	return bt
}

// Equal returns true if the BehaviorTree is equal to the other.
func (bt *BehaviorTree) Equal(other *BehaviorTree) bool {
	if bt.rootNode == other.rootNode {
		return true
	}

	if bt.rootNode.Equal(other.rootNode) {
		return true
	}

	return false
}

// init initializes the BehaviorTree.
func (bt *BehaviorTree) init(str string) {
	for _, line := range strings.Split(str, "\n") {
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "//") {
			continue
		}

		newNode := bt.newNode(line)

		if bt.rootNode == nil {
			bt.rootNode = newNode
		}
	}
}

// newNode creates a new node.
func (bt *BehaviorTree) newNode(str string) Node {
	strs := strings.SplitN(str, ":", 2)

	nodeStr := strings.TrimSpace(strs[0])

	var bodyStr string
	if 1 < len(strs) {
		bodyStr = strings.TrimSpace(strs[1])
	}

	return newNode(nodeStr, bodyStr)

}

// newNode creates a new node.
func newNode(nodeStr, bodyStr string) Node {
	switch nodeStr {
	case "Sequence":
		return &Sequence{}
	case "Selector":
		return &Selector{}
	}

	return nil
}
