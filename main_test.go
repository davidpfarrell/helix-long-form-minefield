package main

import (
	"math"
	"testing"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Common Testing Asserts
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// assertIntEqual
//
func assertIntEqual(t *testing.T, i1 int, i2 int, msg string) {
	if i1 != i2 {
		t.Errorf("%s : '%d' != '%d'", msg, i1, i2)
	}
}

// assertFloat64Equal
//
func assertFloat64Equal(t *testing.T, f1 float64, f2 float64, msg string) {
	if f1 != f2 {
		t.Errorf("%s : '%f' != '%f'", msg, f1, f2)
	}
}

// assertStringEqual
//
func assertStringEqual(t *testing.T, s1 string, s2 string, msg string) {
	if s1 != s2 {
		t.Errorf("%s : '%s' != '%s'", msg, s1, s2)
	}
}

// assertBoolEqual
//
func assertBoolEqual(t *testing.T, b1 bool, b2 bool, msg string) {
	if b1 != b2 {
		t.Errorf("%s : '%t' != '%t'", msg, b1, b2)
	}
}

// assertPanicWithMsg
//
func assertPanicWithMsg(t *testing.T, f func(), s string, msg string) {
	defer func() {
		if r := recover(); r != nil {
			s2 := r.(string)
			if s2 != s {
				t.Errorf("%s : '%s' != '%s'", msg, s, s2)
			}
		} else {
			t.Errorf("%s : Did not panic", msg)
		}
	}()
	f()
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Mine helpers
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Point
//
type Point struct {
	X float64
	Y float64
}

// NewPoint
//
func NewPoint(x, y float64) *Point {
	return &Point{X: x, Y: y}
}

// ToMine
//
func (p *Point) ToMine(power float64) *Mine {
	return NewMine(p.X, p.Y, power)
}

// distance
// Used to back-into a desired distance
//
func distance(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p1.X-p2.X, 2) + math.Pow(p1.Y-p2.Y, 2))
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// TestRunSimulations_SingleResult
//
func TestRunSimulations_SingleResult(t *testing.T) {
	mines := []*Mine{
		NewMine(0, 0, 15),

		NewMine(0, 10, 10),
		NewMine(0, 12, 15),

		NewMine(0, 20, 2),
		NewMine(0, 22, 8),
		NewMine(0, 24, 15),

		NewMine(0, 30, 10),
		NewMine(0, 32, 32),
	}
	results := runSimulations(mines)
	assertIntEqual(t, len(results), 1, "runSimulations: len(results)")
	assertStringEqual(t, results[0].String(), "SimulationResult{ X: 0, Y: 32, Time: 1, Count: 7 }", "runSimulations: results[0]")
}

// TestRunSimulations_MultipleResults
//
func TestRunSimulations_MultipleResults(t *testing.T) {
	mines := []*Mine{
		NewMine(0, 0, 15),

		NewMine(0, 10, 10),
		NewMine(0, 12, 15),

		NewMine(0, 20, 2),
		NewMine(0, 22, 8),
		NewMine(0, 24, 15),

		NewMine(0, 30, 10),
	}
	results := runSimulations(mines)
	assertIntEqual(t, len(results), 2, "runSimulations: len(results)")
	assertStringEqual(t, results[0].String(), "SimulationResult{ X: 0, Y: 12, Time: 1, Count: 5 }", "runSimulations: resuklts[0]")
	assertStringEqual(t, results[1].String(), "SimulationResult{ X: 0, Y: 24, Time: 1, Count: 5 }", "runSimulations: resuklts[0]")
}

// TestRunSimulations_EmptyMines
//
func TestRunSimulations_EmptyMines(t *testing.T) {
	var results []*SimulationResult

	results = runSimulations(nil)
	assertIntEqual(t, len(results), 0, "runSimulations(nil)")

	results = runSimulations([]*Mine{})
	assertIntEqual(t, len(results), 0, "runSimulations(emptyList)")
}
