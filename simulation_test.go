package main

import "testing"

// TestSimulationResult_String
//
func TestSimulationResult_String(t *testing.T) {
	sr := &SimulationResult{Mine: NewMine(1.2, 2.3, 0), Time: 4, Count: 5}
	assertStringEqual(t, sr.String(), "SimulationResult{ X: 1.2, Y: 2.3, Time: 4, Count: 5 }", "SimulationResult.String()")
}

// TestRunSimulation
//
func TestRunSimulation(t *testing.T) {
	mf := NewMinefield()

	// We use mines on a straight line to keep distance logic easy to reason about
	//
	m1 := NewMine(0, 0, 15)

	m2 := NewMine(0, 10, 10)
	m3 := NewMine(0, 12, 15)

	m4 := NewMine(0, 20, 2)
	m5 := NewMine(0, 22, 8)
	m6 := NewMine(0, 24, 15)

	m7 := NewMine(0, 30, 10)
	m8 := NewMine(0, 32, 32)

	mf.Add(m1)
	mf.Add(m2)
	mf.Add(m3)
	mf.Add(m4)
	mf.Add(m5)
	mf.Add(m6)
	mf.Add(m7)
	mf.Add(m8)

	var r *SimulationResult

	r = RunSimulation(mf, m1)
	assertStringEqual(t, r.String(), "SimulationResult{ X: 0, Y: 0, Time: 2, Count: 3 }", "RunSimulation")

	r = RunSimulation(mf, m2)
	assertStringEqual(t, r.String(), "SimulationResult{ X: 0, Y: 10, Time: 1, Count: 3 }", "RunSimulation")

	r = RunSimulation(mf, m3)
	assertStringEqual(t, r.String(), "SimulationResult{ X: 0, Y: 12, Time: 1, Count: 5 }", "RunSimulation")

	r = RunSimulation(mf, m4)
	assertStringEqual(t, r.String(), "SimulationResult{ X: 0, Y: 20, Time: 3, Count: 3 }", "RunSimulation")

	r = RunSimulation(mf, m5)
	assertStringEqual(t, r.String(), "SimulationResult{ X: 0, Y: 22, Time: 1, Count: 3 }", "RunSimulation")

	r = RunSimulation(mf, m6)
	assertStringEqual(t, r.String(), "SimulationResult{ X: 0, Y: 24, Time: 1, Count: 6 }", "RunSimulation")

	r = RunSimulation(mf, m7)
	assertStringEqual(t, r.String(), "SimulationResult{ X: 0, Y: 30, Time: 1, Count: 4 }", "RunSimulation")

	r = RunSimulation(mf, m8)
	assertStringEqual(t, r.String(), "SimulationResult{ X: 0, Y: 32, Time: 1, Count: 7 }", "RunSimulation")
}
