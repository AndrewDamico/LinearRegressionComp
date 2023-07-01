package main

import (
	"fmt"
	"github.com/guptarohit/asciigraph"
	"github.com/montanaflynn/stats"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

// Calculates the mean runtime for each experiment
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

// Generates a table showing speed ratios of each experiment
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

// Generates a performance graph showing runtimes of each experiment
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
