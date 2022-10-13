package main

import (
	"fmt"
	"strconv"
)

// RunSimulation
//
func RunSimulation(minefield *Minefield, startMine *Mine) *SimulationResult {
	// Keep track of all exploded mines
	// Init large map to help enforce  the rule:
	//    "optimize for processing efficiency."
	allExploded := make(map[string]*Mine, minefield.Len()*10)
	// Seed result with starting mine at t=0
	//
	r := &SimulationResult{Mine: startMine, Time: 0, Count: 1}
	allExploded[startMine.Key] = startMine

	for prevExploded, exploded, time := []*Mine{startMine}, []*Mine{}, 1; // Seed prev, t=0 already accounted for
	len(prevExploded) > 0 && minefield.Len()-len(allExploded) > r.Count;  // Stop if no new mines exploded or too few mines left to create new winner
	prevExploded, exploded, time = exploded, []*Mine{}, time+1 {          // Shift new to prev, update time
		// Build list of newly exploded mines based on previously exploded mines
		//
		for _, mine := range prevExploded {
			for _, target := range minefield.GetTargetsFor(mine) {
				// If we haven't exploded this mine previously
				//
				if _, exists := allExploded[target.Key]; !exists {
					// Boom !
					//
					exploded = append(exploded, target)
					allExploded[target.Key] = target
				}
			}
		}
		// Update result if this round produced more newly exploded mines
		// NOTE: Description is not clear when multiple time intervals
		//       (from the same starting mine) have the same explosion count - Punting to first interval
		//
		if len(exploded) > r.Count {
			r.Time = time
			r.Count = len(exploded)
		}
	}
	return r
}

// SimulationResult
// "the mine that, if triggered first, will result in the highest number of explosions occurring during a
//  single time interval."
//
// "Output the coordinates of the winning mine, the time interval of the peak number of explosions,
//  and the number of explosions during that interval."
//
type SimulationResult struct {
	Mine  *Mine // Starting mine
	Time  int   // Time of max explosions
	Count int   // Number of mines exploded
}

// String
// Implements fmt.Stringer
//
func (r *SimulationResult) String() string {
	return fmt.Sprintf(
		"SimulationResult{ X: %s, Y: %s, Time: %d, Count: %d }",
		strconv.FormatFloat(r.Mine.X, 'f', -1, 64),
		strconv.FormatFloat(r.Mine.Y, 'f', -1, 64),
		r.Time,
		r.Count)
}
