package main

import (
	"fmt"
	"math"
	"strconv"
)

// Mine
// We sacrifice immutability, and memoize the key here, to help enforce the rule:
//    "optimize for processing efficiency."
//
type Mine struct {
	X   float64 // X position within minefield
	Y   float64 // Y position within minefield
	P   float64 // Explosive Power
	Key string  // Memoized key
}

// New
// Using New to instantiate Mines to memoize the key and help enforce the rule:
//    "explosive power may not be negative."
//
func NewMine(x, y, p float64) *Mine {
	if p < 0.0 {
		panic("explosive power may not be negative.")
	}
	return &Mine{
		X: x,
		Y: y,
		P: p,
		Key: fmt.Sprintf(
			"%s:%s",
			strconv.FormatFloat(x, 'f', -1, 64),
			strconv.FormatFloat(y, 'f', -1, 64)),
	}
}

// CanReach
// Determine if explosive power is strong enough to reach target mine
// NOTE: Assuming "straight-line distance" to mean a radius from triggered mine
// NOTE: Precision errors abounds - Seems like it would be impossible to automate a test against this challenge
//
func (m *Mine) CanReach(target *Mine) bool {
	// Using A^2 + B^2 = C^2 to compute distance between mines
	//
	return math.Sqrt(math.Pow(m.X-target.X, 2)+math.Pow(m.Y-target.Y, 2)) <= m.P
}

// String
// Implements fmt.Stringer
//
func (m *Mine) String() string {
	return fmt.Sprintf(
		"Mine{ X: %s, Y: %s, P: %s }",
		strconv.FormatFloat(m.X, 'f', -1, 64),
		strconv.FormatFloat(m.Y, 'f', -1, 64),
		strconv.FormatFloat(m.P, 'f', -1, 64))
}
