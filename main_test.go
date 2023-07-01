package main

import (
	"math"
	"testing"
)

// Determines the acceptable tolerance for difference in floating points.
func withinTolerance(a, b, e float64) bool {
	// https://medium.com/pragmatic-programmers/testing-floating-point-numbers-in-go-9872fe6de17f
	if a == b {
		return true
	}
	d := math.Abs(a - b)
	if b == 0 {
		return d < e
	}
	return (d / math.Abs(b)) < e
}

var tolerance = 1e-13 // Best Case

// PYTHON TESTS
// Run on First Anscombe Set
func TestExperimentPythonOne(t *testing.T) {
	t.Parallel()
	var set = "One"
	var got = ExperimentGo(set, 1).Coefficients
	var want = ExperimentPython(set, "1").Coefficients

	// Test Intersect
	if !withinTolerance(want[0], got[0], tolerance) {
		t.Errorf("want %.18f, got %.18f", want, got)
	}
	// Test Slope
	if !withinTolerance(want[1], got[1], tolerance) {
		t.Errorf("want %.18f, got %.18f", want, got)
	}
}

// Run on Second Anscombe Set
func TestExperimentPythonTwo(t *testing.T) {
	t.Parallel()
	var set = "Two"
	var got = ExperimentGo(set, 1).Coefficients
	var want = ExperimentPython(set, "1").Coefficients

	// Test Intersect
	if !withinTolerance(want[0], got[0], tolerance) {
		t.Errorf("want %.18f, got %.18f", want, got)
	}
	// Test Slope
	if !withinTolerance(want[1], got[1], tolerance) {
		t.Errorf("want %.18f, got %.18f", want, got)
	}
}

// Run on Third Anscombe Set
func TestExperimentPythonThree(t *testing.T) {
	t.Parallel()
	var set = "Three"
	var got = ExperimentGo(set, 1).Coefficients
	var want = ExperimentPython(set, "1").Coefficients

	// Test Intersect
	if !withinTolerance(want[0], got[0], tolerance) {
		t.Errorf("want %.18f, got %.18f", want, got)
	}
	// Test Slope
	if !withinTolerance(want[1], got[1], tolerance) {
		t.Errorf("want %.18f, got %.18f", want, got)
	}
}

// Run on Fourth Anscombe Set
func TestExperimentPythonFour(t *testing.T) {
	t.Parallel()
	var set = "Four"
	var got = ExperimentGo(set, 1).Coefficients
	var want = ExperimentPython(set, "1").Coefficients

	// Test Intersect
	if !withinTolerance(want[0], got[0], tolerance) {
		t.Errorf("want %.18f, got %.18f", want, got)
	}
	// Test Slope
	if !withinTolerance(want[1], got[1], tolerance) {
		t.Errorf("want %.18f, got %.18f", want, got)
	}
}

// R TESTS
// Run on First Anscombe Set
func TestExperimentROne(t *testing.T) {
	t.Parallel()
	var set = "One"
	var got = ExperimentGo(set, 1).Coefficients
	var want = ExperimentR(set, "1").Coefficients

	// Test Intersect
	if !withinTolerance(want[0], got[0], tolerance) {
		t.Errorf("want %.18f, got %.18f", want, got)
	}
	// Test Slope
	if !withinTolerance(want[1], got[1], tolerance) {
		t.Errorf("want %.18f, got %.18f", want, got)
	}
}

// Run on Second Anscombe Set
func TestExperimentRTwo(t *testing.T) {
	t.Parallel()
	var set = "Two"
	var got = ExperimentGo(set, 1).Coefficients
	var want = ExperimentR(set, "1").Coefficients

	// Test Intersect
	if !withinTolerance(want[0], got[0], tolerance) {
		t.Errorf("want %.18f, got %.18f", want, got)
	}
	// Test Slope
	if !withinTolerance(want[1], got[1], tolerance) {
		t.Errorf("want %.18f, got %.18f", want, got)
	}
}

// Run on Third Anscombe Set
func TestExperimentRThree(t *testing.T) {
	t.Parallel()
	var set = "Three"
	var got = ExperimentGo(set, 1).Coefficients
	var want = ExperimentR(set, "1").Coefficients

	// Test Intersect
	if !withinTolerance(want[0], got[0], tolerance) {
		t.Errorf("want %.18f, got %.18f", want, got)
	}
	// Test Slope
	if !withinTolerance(want[1], got[1], tolerance) {
		t.Errorf("want %.18f, got %.18f", want, got)
	}
}

// Run on Fourth Anscombe Set
func TestExperimentRFour(t *testing.T) {
	t.Parallel()
	var set = "Four"
	var got = ExperimentGo(set, 1).Coefficients
	var want = ExperimentR(set, "1").Coefficients

	// Test Intersect
	if !withinTolerance(want[0], got[0], tolerance) {
		t.Errorf("want %.18f, got %.18f", want, got)
	}
	// Test Slope
	if !withinTolerance(want[1], got[1], tolerance) {
		t.Errorf("want %.18f, got %.18f", want, got)
	}
}

/*
func TestExperimentGo(t *testing.T) {
	t.Parallel()
	var response Response = ExperimentGo("One", 1)
	target := ExperimentPython("One", 1)
	target := []float64{3.0000909090909085, 0.5000909090909091}
	want := target[0]
	var got = response.Coefficients[0]

	if !withinTolerance(want, got, 1e-12) {
		t.Errorf("want %.18f, got %.18f", want, got)
	}
}

*/
