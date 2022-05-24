package behaviortree

import "log"

// Sequence represents a BehaviorTree sequence.
//
// Sequence is a composite node that runs its children sequentially until one of them returns a failure.
type Sequence struct {
	children []Node
}

// Equal returns true if the Sequence is equal to the other.
func (s *Sequence) Equal(other Node) bool {
	log.Printf("Unimplemented Equal() for children")

	_, ok := other.(*Sequence)
	if !ok {
		return false
	}

	return true
}

func (s *Sequence) Tick() State {
	for _, child := range s.children {
		state := child.Tick()

		if state != Success {
			return state
		}
	}

	return Success
}
