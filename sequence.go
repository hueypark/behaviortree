package behaviortree

// Sequence represents a BehaviorTree sequence.
//
// Sequence is a composite node that runs its children sequentially until one of them returns a failure.
type Sequence struct {
	children []Node
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
