package main

import "github.com/montanaflynn/stats"

// Create function to convert slices to coordinates for Go regression package.
func makeCoordinates(x, y []float64) []stats.Coordinate {
	// Converts two slices containing x and y values returning a set of coordinates
	container := make([]stats.Coordinate, len(x))

	for i := 0; i < len(x); i++ {
		container[i] = stats.Coordinate{
			x[i], y[i],
		}
	}
	return container
}
