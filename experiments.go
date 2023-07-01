package main

import (
	"encoding/json"
	"fmt"
	"github.com/montanaflynn/stats"
	"os/exec"
	"time"
)

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