package gosearch

import "strconv"

// State is a basic state represtation for search algorithms
type State interface {

	// For and input action, the function generates a new state with
	// the new action included
	ApplyAction(action Action) State

	// Returns the list of actions applied to the state to research the
	// actual state
	GetPartialSolution() []Action

	// Returns the sum of all the costs for all the actions applied
	// to the actual state
	GetSolutionCost() float64

	// For a given action, the funcion determinate if is possible to
	// apply that action to the state
	//	isValidAction(action Action) bool

	// Method that generate a list of all the possible applicable actions
	// for the current state
	GetApplicableActions() []Action

	// Returns if the actual state is a solution state
	IsSolution() bool

	// Compare two states
	Equal(second State) bool

	// Add the action to the current state
	//	addActionToSolution(action Action)

	// Returns the depth in the search tree of the current state
	GetStateLevel() int

	// Default string representation (mainly for debug)
	String() string
}

// Action interface to represent the cost of an action
type Action interface {
	// represents the float cost for an Action
	Cost() float64
}

// Heuristic interface to represente the heuristic value for a state
type Heuristic interface {
	// the heuristic evaluation for a state
	Heuristic() float64
}

// Statistics information about the state space explored by the search
type Statistics struct {
	NodesExplored   int
	NodesDuplicated int
	MaxDepth        int
	Solutions       int
}

// Basic string default representation for the Statistics
func (stats Statistics) String() string {

	return "[" +
		"NodesExplored: " + strconv.Itoa(stats.NodesExplored) + ", " +
		"NodesDuplicated: " + strconv.Itoa(stats.NodesDuplicated) + ", " +
		"MaxDepth: " + strconv.Itoa(stats.MaxDepth) + ", " +
		"Solutions: " + strconv.Itoa(stats.Solutions) +
		"]"

}

// Search mechanism

// SearchBreadthFirst is a basic search without domain information
// BreadthFirst search algorithm
// (https://en.wikipedia.org/wiki/Breadth-first_search) to search the
// solution for a initial state provided.  The initial state of the
// problem must be provided and as result the algorithm returns the
// list of solution action (if the problem as solution) and a basic
// statistics about the nodes explored, duplicate nodes and the
// maximum depth explored.
func SearchBreadthFirst(initialState State) ([]Action, Statistics) {

	return findFirstSolution(initialState, new(queue))
}

// SearchDepthFirst is a basic search without domain information Depth
// search algorithm (https://en.wikipedia.org/wiki/Depth-first_search)
// to search the solution for a initial state provided.  The initial
// state of the problem must be provided and as result the algorithm
// returns the list of solution action (if the problem as solution)
// and a basic statistics about the nodes explored, duplicate nodes
// and the maximum depth explored.
func SearchDepthFirst(initialState State) ([]Action, Statistics) {

	return findFirstSolution(initialState, new(stack))
}

// SearchIterativeDepth is a basic search without domain information
// Iterative Depth search algorithm
// (https://en.wikipedia.org/wiki/Iterative_deepening_depth-first_search)
// to search the solution for a initial state provided.  For each
// iteration, the depth in the search is incremented in 1 level.  The
// initial state of the problem must be provided and as result the
// algorithm returns the list of solution action (if the problem as
// solution) and a basic statistics about the nodes explored,
// duplicate nodes and the maximum depth explored.
func SearchIterativeDepth(initial State) ([]Action, Statistics) {

	// linear incremental
	var solution []Action = []Action{}
	var maxDepth int
	stats := Statistics{NodesExplored: 0, NodesDuplicated: 0, MaxDepth: 0, Solutions: 0}
	var statistics Statistics
	depth := 1

	for len(solution) == 0 {
		solution, maxDepth, statistics =
			findFirstSolutionAux(initial, new(stack), depth)
		// aggregate stats
		stats.NodesExplored += statistics.NodesExplored
		stats.NodesDuplicated += statistics.NodesDuplicated
		stats.MaxDepth = max(stats.MaxDepth, maxDepth)
		stats.Solutions += statistics.Solutions
		if depth > maxDepth {
			return []Action{}, stats // no solution
		}
		depth++
	}

	return solution, stats
}

// SearchAstar implement an Astar algorithm
// (https://en.wikipedia.org/wiki/A*_search_algorithm) to search a
// solution state for a problem. The State must implement also the
// Heuristic interface.
// The initial state of the problem must be provided and as result
// the algorithm returns the list of solution action (if the problem
// as solution) and a basic statistics about the nodes explored,
// duplicate nodes and the maximum depth explored.
func SearchAstar(initialState State) ([]Action, Statistics) {

	return findFirstSolutionAstar(initialState, new(floatPriorityList))
}
