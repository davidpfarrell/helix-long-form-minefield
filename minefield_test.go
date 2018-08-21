package main

import (
	"fmt"
	"math"
	"testing"
)

// TestMinefield
//
func TestMinefield(t *testing.T) {
	mf := NewMinefield()

	m1 := NewMine(0, 0, 1) // 1 == too short

	mf.Add(m1)
	assertIntEqual(t, mf.Len(), 1, "Minefield.Len()")

	m2 := NewMine(1, 1, math.Sqrt(2)) // Distance == sqrt(2)
	mf.Add(m2)
	assertIntEqual(t, mf.Len(), 2, "Minefield.Len()")

	t1 := mf.GetTargetsFor(m1)
	assertIntEqual(t, len(t1), 0, "Minefield.GetTargetsFor()")

	t2 := mf.GetTargetsFor(m2)
	assertIntEqual(t, len(t2), 1, "Minefield.GetTargetsFor()")
	assertBoolEqual(t, t2[0] == m1, true, "Mine2 target[0] = Mine1")
}

// TestNewMinefieldDuplicate
//
func TestNewMinefieldDuplicate(t *testing.T) {
	mf := NewMinefield()

	m1 := NewMine(0, 0, 1)
	mf.Add(m1)
	assertIntEqual(t, mf.Len(), 1, "Minefield.Len()")

	m2 := NewMine(0, 0, 1)
	f := func() {
		mf.Add(m2)
	}

	assertPanicWithMsg(
		t,
		f,
		fmt.Sprintf("There may not be more than one mine at the same coordinates: %s : %s", m1, m2),
		"Minefield: Add mine with duplicate coordinates")
}
