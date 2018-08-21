package main

import "fmt"

// Minefield
// Maintains a map in order to enforce the rules:
//    "There may not be more than one mine at the same coordinates."
//    "optimize for processing efficiency."
//
type Minefield struct {
	mines map[string]*node
}

// NewMinefield
//
func NewMinefield() *Minefield {
	mf := &Minefield{mines: map[string]*node{}}
	return mf
}

// NewMinefieldWithMines
//
func NewMinefieldWithMines(mines []*Mine) *Minefield {
	mf := NewMinefield()
	for _, mine := range mines {
		mf.Add(mine)
	}
	return mf
}

// Add
//
func (mf *Minefield) Add(mine *Mine) {
	if existingNode, exists := mf.mines[mine.Key]; exists {
		panic(fmt.Sprintf("There may not be more than one mine at the same coordinates: %s : %s", existingNode.mine, mine))
	}
	// Initiate as nil to reduce memory in case we don't add any targets
	//
	var targets []*Mine = nil
	// Update targets list for mines within range
	// NOTE: I like the canReach abstraction, but it does require two calls.
	//       One could imagine extracting it and only calling it once here.
	// TODO This doesn't scale well and may break the entire test
	//
	for _, node := range mf.mines {
		// Is new mine within range of existing mine?
		//
		if node.mine.CanReach(mine) {
			node.targets = append(node.targets, mine)
		}
		// Is existing mine within range of new mine?
		//
		if mine.CanReach(node.mine) {
			targets = append(targets, node.mine)
		}
	}
	mf.mines[mine.Key] = &node{mine: mine, targets: targets}
}

// Len
//
func (mf *Minefield) Len() int {
	return len(mf.mines)
}

// GetTargetsFor
// List of mines that are within blast range of the specified mine
//
func (mf *Minefield) GetTargetsFor(mine *Mine) []*Mine {
	if node, exists := mf.mines[mine.Key]; exists && node.mine == mine {
		return node.targets
	}
	panic(fmt.Sprintf("Mine not found in minefiend: %s", mine))
}

// node
// Stores a slice of mines that would be triggered if the parent mine went off.
// We compute this graph once to enforce the rule:
//    "optimize for processing efficiency."
//
type node struct {
	mine    *Mine
	targets []*Mine
}
