package main

import (
	"encoding/json"
	"fmt"
	"github.com/guptarohit/asciigraph"
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
var nRuns int = 100
var nExperiments = 20

// Run Go Experiment
func ExperimentGo(set string, nRuns int) Response {
	//fmt.Println("  Performing Go Experiment...")

	var responseGo Response
	var times []float64

	for i := 0; i < nRuns; i++ {
		startTime := time.Now()
		points, _ := stats.LinearRegression(
			MakeCoordinates(
				anscombe[set]["x"],
				anscombe[set]["y"]),
		)
		responseGo.Coefficients = EquationLine(points)
		elapsedTime := time.Since(startTime).Seconds()
		elapsedTime = float64(elapsedTime)
		times = append(times, elapsedTime)
	}

	responseGo.Time, _ = stats.Mean(times)

	return responseGo
}

// Run Python Experiment
func ExperimentPython(set, nRunsString string) Response {
	//fmt.Println("  Performing Python Experiment...")
	var responsePython Response

	// Run the Python Script
	args := []string{"AnscombeTest.py", set, nRunsString}
	cmd := exec.Command("python", args...)
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
func ExperimentR(set, nRunsString string) Response {
	//fmt.Println("  Performing R Experiment...")
	var responseR Response

	// Run the R Script
	args := []string{"AnscombeTest.R", set, nRunsString}
	cmd := exec.Command("Rscript", args...)
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
		fmt.Printf("Performing Experiment on set %s: %d \r", set, i+1)
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
	// #ToDo make averages of experiment results instead
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

		_, err = fmt.Scanf("%d", &choice)
		if err != nil {
			choice = -1
		}

		switch choice {
		case 1:
			fmt.Println("Running all Four Anscombe Quartets.")
			experiment("One")
			experiment("Two")
			experiment("Three")
			experiment("Four")
		case 2:
			mainMenu()
		default:
			mainMenu()
		}

		//if choice == 0 {
		return
	}
}

func performanceGraph() {

	if len(performanceGo) == 0 {
		fmt.Println("No experiments have been run in this session. Please run some experiments first.")
		runall()
	}
	log := [][]float64{
		performanceGo,
		performanceR,
		performancePython,
	}

	graph := asciigraph.PlotMany(log, asciigraph.Precision(10), asciigraph.SeriesColors(
		asciigraph.Red,
		asciigraph.Yellow,
		asciigraph.Green,
		//asciigraph.Blue,
	))

	fmt.Println("Legend: Go (Red), Python (Green), R (Yellow)")
	fmt.Println(graph)
	fmt.Println()
}
func calcPerformance() {
	meanPython, _ := stats.Mean(performancePython)
	meanR, _ := stats.Mean(performanceR)
	meanGo, _ := stats.Mean(performanceGo)
	fmt.Println()
	fmt.Println("Mean Python Runtime:", meanPython)
	fmt.Println("Mean R Runtime:", meanR)
	fmt.Println("Mean Go Runtime:", fmt.Sprintf("%.10f", meanGo))
	fmt.Println()
	performanceMatrix(meanPython, meanR, meanGo)
}
func performanceMatrix(meanPython, meanR, meanGo float64) {
	// Check to make sure that experiments exist. If not, give option to run all experiments or return to menu.
	if len(performanceGo) == 0 {
		fmt.Println("No experiments have been run in this session. Please run some experiments first.")
		runall()
		calcPerformance()
		return
	}
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
	if len(performanceGo) == 0 {
		fmt.Println()
		fmt.Println("No experiments have been run in this session. Please run some experiments first.")
		return
	}
	fmt.Println()

	return
}

func main() {
	fmt.Println("Regression Performance between Python, R, and Go")
	fmt.Println()
	mainMenu()
}
