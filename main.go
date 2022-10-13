package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"sort"
	"sync"
)

/*

Write a program which takes as input a list of mines composing a 2D minefield;
each mine has an X position, a Y position, and an explosive power.
All three parameters may be assumed to be single-precision floats;
explosive power may not be negative.
There may not be more than one mine at the same coordinates.

When a mine in the minefield is triggered at time T=0,
it causes all other mines within a straight-line distance
less than or equal to its explosive power to be triggered at time T=1.

Those mines subsequently trigger additional mines at T=2, and soforth, in a chain reaction.

Have your program determine, for any given input minefield,
the mine that, if triggered first, will result in the highest number of explosions occurring during a single time interval.

Output the coordinates of the winning mine, the time interval of the peak number of explosions,
and the number of explosions during that interval.

In case of a tie, output each of the best mines, sorted by X coordinate then Y coordinate.

Assume that the minefield may be large, but not larger than can easily fit in memory;
optimize for processing efficiency.

*/

// main
//
func main() {
	mines := readMines(os.Stdin)
	fmt.Printf("Minefield: %v\n", mines)
	// Empty minefield?
	if len(mines) == 0 {
		fmt.Println("Minefield is empty")
	} else {
		// Let's asplode some mines !
		//
		results := runSimulations(mines)
		// Print results
		//
		for _, result := range results {
			// With no specifics on output format, serialized result should be sufficient
			//
			fmt.Println(result)
		}
	}
}

// readMines
//
func readMines(in io.Reader) []*Mine {
	// Read in into slice of Mines
	//
	mines := []*Mine{}
	scanner := bufio.NewScanner(in)
	var quit bool = false
	for !quit && scanner.Scan() {
		var x, y, p float64
		// Expect full line, quit reading on first failure
		//
		if n, err := fmt.Sscanf(scanner.Text(), "%f %f %f", &x, &y, &p); err == nil && n == 3 {
			mines = append(mines, NewMine(x, y, p))
		} else {
			quit = true
		}
	}
	return mines
}

// runSimulations
// Run simulation against every mine as starting mine
// Should tolerate nil/empty_list as input
//
func runSimulations(mines []*Mine) []*SimulationResult {
	// Initate to empty list, will append on first result
	//
	results := []*SimulationResult{}
	// Only need to create the minefield once
	//
	minefield := NewMinefieldWithMines(mines)
	// n < 1 == Query without change
	//
	maxProcs := runtime.GOMAXPROCS(0)
	// Tracks access to mines withing go routines
	//
	index := 0
	// Guards access to index
	//
	mutex := &sync.Mutex{}
	// Used for go routines to return results
	//
	ch := make(chan *SimulationResult)
	// Try to optimize CPU by generating a go routine for each processor,
	// or for each mine if less mines than maxProcs
	//
	for p := 0; p < maxProcs && p < len(mines); p++ { // p is unused
		go func() {
			for {
				// Get index of next mine to process
				//
				mutex.Lock()
				i := index
				index++
				mutex.Unlock()
				// Anything left to process?
				//
				if i >= len(mines) {
					return
				}
				// Run the simulation and return the result
				//
				ch <- RunSimulation(minefield, mines[i])
			}
		}()
	}
	// Fetch results from go routines
	//
	for i := 0; i < len(mines); i++ { // i is unused
		// Fetch result
		//
		result := <-ch
		// Should we keep it?
		//
		if len(results) == 0 || results[0].Count == result.Count { // First check will fall through to append
			// First result, or tie for best results
			//
			results = append(results, result)
		} else if results[0].Count < result.Count {
			// New best result
			//
			results = []*SimulationResult{result}
		}
	}
	// Sort results (if needbe)
	//    "In case of a tie, output each of the best mines, sorted by X coordinate then Y coordinate."
	//
	if len(results) > 1 {
		sort.Slice(results, func(i, j int) bool {
			return results[i].Mine.X < results[j].Mine.X || (results[i].Mine.X == results[j].Mine.X && results[i].Mine.Y < results[j].Mine.Y)
		})
	}
	return results
}
