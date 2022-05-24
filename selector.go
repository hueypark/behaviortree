package behaviortree

// Selector represents a BehaviorTree selector.
//
// Selector is a composite node that runs its children sequentially until one of them returns a success.
type Selector struct {
	children []Node
}

func Tick(s *Selector) State {
	for _, child := range s.children {
		state := child.Tick()

		if state == Success {
			return state
		}
	}

	return Failure
}
