package main

import "github.com/montanaflynn/stats"

// Create function to determine equation of line
func EquationLine(points []stats.Coordinate) []float64 {

	// determine min and max coordinates from a set of coordinates
	coords := MinMaxCoordinates(points)

	x1 := coords[0].X
	y1 := coords[0].Y
	x2 := coords[len(coords)-1].X
	y2 := coords[len(coords)-1].Y

	// Calculate slope and intercept using y = mx + b
	m := (y2 - y1) / (x2 - x1)
	b := y1 - m*x1

	container := []float64{b, m}

	return container
}

// Create function to determine the minimum and maximum coordinates of a set of coordinates
func MinMaxCoordinates(x []stats.Coordinate) []stats.Coordinate {
	minX := x[0].X
	maxX := x[0].X
	minXY := x[0]
	maxXY := x[0]

	for _, point := range x {
		if point.X < minX {
			minX = point.X
			minXY = point
		}
		if point.X > maxX {
			maxX = point.X
			maxXY = point
		}
	}
	container := []stats.Coordinate{minXY, maxXY}

	return container
}
