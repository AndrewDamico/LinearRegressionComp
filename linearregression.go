package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
	"os/exec"
	"time"
)

func LinearRegression(s stats.Series) (gradient, intercept float64, err error) {
	//edited from the montanaflynn/stats package: https://github.com/montanaflynn/stats/blob/master/regression.go

	// Placeholder for the math to be done
	var sum [5]float64

	// Loop over data keeping index in place
	i := 0
	for ; i < len(s); i++ {
		sum[0] += s[i].X
		sum[1] += s[i].Y
		sum[2] += s[i].X * s[i].X
		sum[3] += s[i].X * s[i].Y
		sum[4] += s[i].Y * s[i].Y
	}

	// Find gradient and intercept
	f := float64(i)
	gradient = (f*sum[3] - sum[0]*sum[1]) / (f*sum[2] - sum[0]*sum[0])
	intercept = (sum[1] / f) - (gradient * sum[0] / f)

	return gradient, intercept, nil
}

// create the Anscomb data sets
var anscombe = map[string]map[string][]float64{
	"One": map[string][]float64{
		"x": []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		"y": []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
	},
	"Two": map[string][]float64{
		"x": []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		"y": []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74},
	},
	"Three": map[string][]float64{
		"x": []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		"y": []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
	},
	"Four": map[string][]float64{
		"x": []float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8},
		"y": []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89},
	},
}

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

// create function to perform linear regression
func model(x []stats.Coordinate) []float64 {
	// Takes a set of coordinates and returns a regression line.
	g, i, _ := LinearRegression(x)
	//fmt.Println(r)
	container := []float64{i, g}

	return container
}

// create the Anscomb data sets

func main() {
	//Run the Go Script
	startTime := time.Now()
	set := "Three"
	x := model(makeCoordinates(anscombe[set]["x"], anscombe[set]["y"]))
	elapsedTime := time.Since(startTime)
	fmt.Println(x)
	fmt.Printf("Elapsed time: %.9f seconds\n", elapsedTime.Seconds())
	// Run the Python Script
	cmd := exec.Command("python", "Anscombe_test.py", set)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(output))
}
