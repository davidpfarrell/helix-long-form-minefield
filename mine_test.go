package main

import (
	"testing"
)

// TestMineZero
//
func TestMineZero(t *testing.T) {
	m := NewMine(0, 0, 0)
	assertFloat64Equal(t, m.X, 0, "Mine.X")
	assertFloat64Equal(t, m.Y, 0, "Mine.Y")
	assertFloat64Equal(t, m.P, 0, "Mine.P")
	assertStringEqual(t, m.Key, "0:0", "Mine.Key")
	assertStringEqual(t, m.String(), "Mine{ X: 0, Y: 0, P: 0 }", "Mine.String()")
}

// TestMinePositive
//
func TestMinePositive(t *testing.T) {
	m := NewMine(1.2, 2.3, 3.4)
	assertFloat64Equal(t, m.X, 1.2, "Mine.X")
	assertFloat64Equal(t, m.Y, 2.3, "Mine.Y")
	assertFloat64Equal(t, m.P, 3.4, "Mine.P")
	assertStringEqual(t, m.Key, "1.2:2.3", "Mine.Key")
	assertStringEqual(t, m.String(), "Mine{ X: 1.2, Y: 2.3, P: 3.4 }", "Mine.String()")
}

// TestMineNegative
//
func TestMineNegative(t *testing.T) {
	m := NewMine(-1.2, -2.3, 3.4)
	assertFloat64Equal(t, m.X, -1.2, "Mine.X")
	assertFloat64Equal(t, m.Y, -2.3, "Mine.Y")
	assertFloat64Equal(t, m.P, 3.4, "Mine.P")
	assertStringEqual(t, m.Key, "-1.2:-2.3", "Mine.Key")
	assertStringEqual(t, m.String(), "Mine{ X: -1.2, Y: -2.3, P: 3.4 }", "Mine.String()")
}

// TestMineNegativePower
//
func TestMineNegativePower(t *testing.T) {
	f := func() {
		NewMine(0, 0, -1)
	}
	assertPanicWithMsg(t, f, "explosive power may not be negative.", "Mine with negative power")
}

// TestMineCanReachTrue
//
func TestMineCanReachTrue(t *testing.T) {
	p1 := NewPoint(-5.4321, -6.5432)
	p2 := NewPoint(7.6454, 8.7654)

	delta := distance(p1, p2)

	m1 := p1.ToMine(delta)
	m2 := p2.ToMine(delta)

	assertBoolEqual(t, m1.CanReach(m2), true, "Mine1.CanReach(Mine2)")
	assertBoolEqual(t, m1.CanReach(m2), true, "Mine2.CanReach(Mine1")
}

// TestMineCanReachFalse
//
func TestMineCanReachFalse(t *testing.T) {
	p1 := NewPoint(-5.4321, -6.5432)
	p2 := NewPoint(7.6454, 8.7654)

	delta := distance(p1, p2)

	// Set power *just* out of range
	//
	m1 := p1.ToMine(delta - 0.001)
	m2 := p2.ToMine(delta - 0.001)

	assertBoolEqual(t, m1.CanReach(m2), false, "Mine1.CanReach(Mine2)")
	assertBoolEqual(t, m1.CanReach(m2), false, "Mine2.CanReach(Mine1)")
}
