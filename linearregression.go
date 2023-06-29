package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/montanaflynn/stats"
	"os"
	"os/exec"
	"time"
)

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

// Create function to determine equation of line
func line(points []stats.Coordinate) []float64 {

	// determine min and max coordinates from a set of coordinates
	coords := minmaxCoordinates(points)

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
func minmaxCoordinates(x []stats.Coordinate) []stats.Coordinate {
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

// create function to perform linear regression
func model(x []stats.Coordinate) []float64 {
	// Takes a set of coordinates and returns a regression line.
	points, _ := stats.LinearRegression(x)

	// calculate the intercept and slope of the points
	container := line(points)

	return container
}

// Define the expected response type from Python and R scripts
type Response struct {
	Line []float64 `json:"line"`
	Time float64   `json:"time"`
}

// Run Experiment
func experiment(set string) {
	var times []float64
	var coefficients []float64

	// Calculate coefficients in Go and run experiment n times
	n := 10
	for i := 0; i < n; i++ {
		startTime := time.Now()
		coefficients = model(makeCoordinates(anscombe[set]["x"], anscombe[set]["y"]))
		elapsedTime := time.Since(startTime).Seconds()
		elapsedTime = float64(elapsedTime)
		times = append(times, elapsedTime)
	}

	averageTime, _ := stats.Mean(times)

	// Run the Python Script
	cmd := exec.Command("python", "Anscombe_test.py", set)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Parse the Python response
	var response Response
	err2 := json.Unmarshal([]byte(output), &response)
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}

	// return statistics
	fmt.Println("---------------")
	fmt.Println("Set:", set)
	fmt.Println("Intercept & Slope (GO):", coefficients)
	fmt.Println("Intercept & Slope (Python):", response.Line)

	fmt.Printf("Elapsed time (GO): %.9f seconds\n", averageTime)
	fmt.Printf("Elapsed time (Python): %.9f seconds\n", response.Time)
	fmt.Println("---------------")
}

func main() {
	fmt.Println("Regression Performance between Python, R, and Go")
	fmt.Println()

	// determine which set to test on

	var choice int64 = -1
	var set string

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Regression Performance - Main Menu")
		fmt.Println("Please choose one of the following options:")

		fmt.Println("1. Calculate Performance Using Anselm Set 1")
		fmt.Println("2. Calculate Performance Using Anselm Set 2")
		fmt.Println("3. Calculate Performance Using Anselm Set 3")
		fmt.Println("4. Calculate Performance Using Anselm Set 4")
		fmt.Println("0. Exit")

		var err error

		_, err = fmt.Scanf("%d", &choice)
		if err != nil {
			choice = -1
		}

		switch choice {
		case 0:
			// Exit the program
			fmt.Println("Goodbye...")
		case 1:
			set = "One"
			fmt.Println("Performing Analysis:")
			experiment(set)
		case 2:
			set = "Two"
			fmt.Println("Performing Analysis:")
			experiment(set)
		case 3:
			set = "Three"
			fmt.Println("Performing Analysis:")
			experiment(set)
		case 4:
			set = "Four"
			fmt.Println("Performing Analysis:")
			experiment(set)
		default:
			fmt.Println("Invalid choice! Please try again.")
		}

		if choice == 0 {
			return
		}
		fmt.Println("Press Enter to continue...")
		scanner.Scan() // Wait for user to press Enter

	}
}
