package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/montanaflynn/stats"
	"github.com/olekukonko/tablewriter"
	"os"
	"os/exec"
	"strconv"
	"time"
)

// Define the expected response type from each experiment
type Response struct {
	Coefficients []float64 `json:"coefficients"`
	Time         float64   `json:"time"`
}

// Create empty containers for the runtime of each experiment
var performancePython []float64
var performanceGo []float64
var performanceR []float64

// Run Go Experiment
func experimentGo(set string) Response {
	fmt.Println("  Performing Go Experiment...")

	var responseGo Response
	var times []float64

	n := 10
	for i := 0; i < n; i++ {
		startTime := time.Now()
		points, _ := stats.LinearRegression(
			makeCoordinates(
				anscombe[set]["x"],
				anscombe[set]["y"]),
		)
		responseGo.Coefficients = equationLine(points)
		elapsedTime := time.Since(startTime).Seconds()
		elapsedTime = float64(elapsedTime)
		times = append(times, elapsedTime)
	}

	responseGo.Time, _ = stats.Mean(times)

	return responseGo
}

// Run Python Experiment
func experimentPython(set string) Response {
	fmt.Println("  Performing Python Experiment...")
	var responsePython Response

	// Run the Python Script
	cmd := exec.Command("python", "AnscombeTest.py", set)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error Python:", err)
	}

	// parse response
	err = json.Unmarshal([]byte(output), &responsePython)
	if err != nil {
		fmt.Println("Error Python Response:", err)
	}

	return responsePython
}

// Run R Experiment
func experimentR(set string) Response {
	fmt.Println("  Performing R Experiment...")
	var responseR Response

	// Run the R Script
	cmd := exec.Command("Rscript", "AnscombeTest.R", set)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error R:", err)
	}
	// Parse response
	err = json.Unmarshal([]byte(output), &responseR)
	if err != nil {
		fmt.Println("Error R Response:", err)
	}

	return responseR
}

// Create table from results of each experiment
func createTable(resultsGo, resultsPython, resultsR Response) {
	// Create Table
	data := [][]interface{}{
		{"Go", resultsGo.Coefficients[0], resultsGo.Coefficients[1], resultsGo.Time},
		{"Python", resultsPython.Coefficients[0], resultsPython.Coefficients[1], resultsPython.Time},
		{"R", resultsR.Coefficients[0], resultsR.Coefficients[1], resultsR.Time},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Language", "Intersect", "Slope", "Runtime"})

	table.SetAutoWrapText(false)

	for _, row := range data {
		strRow := make([]string, len(row))
		for i, val := range row {
			if floatValue, ok := val.(float64); ok {
				strRow[i] = strconv.FormatFloat(floatValue, 'f', -1, 64)
			} else {
				strRow[i] = fmt.Sprint(val)
			}
		}
		table.SetBorder(false)
		table.Append(strRow)
	}

	table.Render()
	fmt.Println()

	return
}

// Run Experiment
func experiment(set string) {
	fmt.Println("Performing Analysis:")

	responseGo := experimentGo(set)
	responsePython := experimentPython(set)
	responseR := experimentR(set)

	// Save Performance Times
	performancePython = append(performancePython, responsePython.Time)
	performanceGo = append(performanceGo, responseGo.Time)
	performanceR = append(performanceR, responseR.Time)

	fmt.Println()
	createTable(responseGo, responsePython, responseR)

	return
}

func performanceMatrix(meanPython, meanR, meanGo float64) {
	// Creates a table showing the speed of execution relative to a baseline language.
	data := [][]interface{}{
		{"Go", meanGo / meanGo, meanPython / meanGo, meanR / meanGo},
		{"Python", meanGo / meanPython, meanPython / meanPython, meanR / meanPython},
		{"R", meanGo / meanR, meanPython / meanR, meanR / meanR},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Baseline", "Go (Numerator) ", "Python (Numerator)", "R (Numerator)"})

	table.SetAutoWrapText(false)

	for _, row := range data {
		strRow := make([]string, len(row))
		for i, val := range row {
			if floatValue, ok := val.(float64); ok {
				strRow[i] = strconv.FormatFloat(floatValue, 'f', -1, 64)
			} else {
				strRow[i] = fmt.Sprint(val)
			}
		}
		table.SetBorder(false)
		table.Append(strRow)
	}

	table.Render()
	fmt.Println()

	return
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
		fmt.Println("5. Calculate Average Performance for All Tests in Current Session")
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
			experiment(set)
		case 2:
			set = "Two"
			experiment(set)
		case 3:
			set = "Three"
			experiment(set)
		case 4:
			set = "Four"
			experiment(set)
		case 5:
			fmt.Println("Calculating Performance on all runs:")
			meanPython, _ := stats.Mean(performancePython)
			meanR, _ := stats.Mean(performanceR)
			meanGo, _ := stats.Mean(performanceGo)
			fmt.Println()
			fmt.Println("Mean Python Runtime:", meanPython)
			fmt.Println("Mean R Runtime:", meanR)
			fmt.Println("Mean Go Runtime:", fmt.Sprintf("%.10f", meanGo))
			fmt.Println()
			performanceMatrix(meanPython, meanR, meanGo)
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
