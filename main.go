package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
	"strconv"
)

//TODO Add performance logs for windows vs linux

// Define the expected response type from each experiment
type Response struct {
	Coefficients []float64 `json:"coefficients"`
	Time         float64   `json:"time"`
}

// Create empty containers for the runtime of each experiment
var performancePython []float64
var performanceGo []float64
var performanceR []float64
var nRuns int = 500 //!important for Windows
var nExperiments = 15
var roundCoefficients uint = 7
var roundTime uint = 8

// Run Experiment
func experiment(set string) {
	fmt.Println("Performing Analysis:")
	fmt.Println("  Number of runs for each test:", nRuns)
	fmt.Println("  Number of runs for each experiment:", nExperiments)
	fmt.Println("  Total number of runs:", nExperiments*nRuns)

	var nRunsString string = strconv.Itoa(nRuns)

	var performanceGoExperiment []float64
	var performancePythonExperiment []float64
	var performanceRExperiment []float64

	var responseGo Response
	var responsePython Response
	var responseR Response

	// Runs each experiment n times and saves experiment values
	for i := 0; i < nExperiments; i++ {
		fmt.Printf("  Performing Experiment on set %s: %d / %d \r", set, i+1, nExperiments)
		responseGo = ExperimentGo(set, nRuns)
		responsePython = ExperimentPython(set, nRunsString)
		responseR = ExperimentR(set, nRunsString)

		// Saves each experiment to the global log for charting
		performanceGoExperiment = append(performanceGoExperiment, responseGo.Time)
		performancePythonExperiment = append(performancePythonExperiment, responsePython.Time)
		performanceRExperiment = append(performanceRExperiment, responseR.Time)
	}

	// Save Performance Times
	performancePython = append(performancePython, performancePythonExperiment...)
	performanceGo = append(performanceGo, performanceGoExperiment...)
	performanceR = append(performanceR, performanceRExperiment...)
	fmt.Println()
	fmt.Println("  Performance times logged.")

	// Displays most recent experiment results
	responseGo.Time, _ = stats.Mean(performanceGoExperiment)
	responseR.Time, _ = stats.Mean(performanceRExperiment)
	responsePython.Time, _ = stats.Mean(performancePythonExperiment)
	fmt.Println()
	fmt.Printf("Results from Set %s: \r", set)
	fmt.Println()
	createTable(responseGo, responsePython, responseR)

	return
}

func runall() {
	var choice int64 = -1
	//scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Would you like to run all Anscombe Quartet sets now?")
		fmt.Println("1. Yes")
		fmt.Println("2. No, Return to Menu")

		var err error

		_, err = fmt.Scanf("%d\n", &choice)
		if err != nil {
			choice = -1
		}

		switch choice {
		case 1:
			clearScreen()
			fmt.Println("Running all Four Anscombe Quartets.")
			experiment("One")
			experiment("Two")
			experiment("Three")
			experiment("Four")
		case 2:
			clearScreen()
			mainMenu()
		default:
			clearScreen()
			mainMenu()
		}

		//if choice == 0 {
		return
	}
}

func main() {
	fmt.Println("Regression Performance between Python, R, and Go")
	fmt.Println()
	mainMenu()
}
