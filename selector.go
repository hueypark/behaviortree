package behaviortree

import "log"

// Selector represents a BehaviorTree selector.
//
// Selector is a composite node that runs its children sequentially until one of them returns a success.
type Selector struct {
	children []Node
}

// Equal returns true if the Selector is equal to the other.
func (s *Selector) Equal(node Node) bool {
	log.Printf("Unimplemented Equal()")

	return false
}

func (s *Selector) Tick() State {
	for _, child := range s.children {
		state := child.Tick()

		if state == Success {
			return state
		}
	}

	return Failure
}
