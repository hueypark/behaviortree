package behaviortree

// State is a state of a node.
type State int

const (
	Success State = iota
	Failure
	Running
)

// Node represents a node in a BehaviorTree.
type Node interface {
	Equal(other Node) bool
	Tick() State
}
