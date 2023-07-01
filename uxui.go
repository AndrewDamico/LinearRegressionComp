package main

import (
	"bufio"
	"fmt"
	"github.com/inancgumus/screen"
	"os"
)

func clearScreen() {
	screen.Clear()
	screen.MoveTopLeft()
}

func mainMenu() {
	clearScreen()
	var choice int64 = -1
	var set string

	scanner := bufio.NewScanner(os.Stdin)

	menuReset := func() {
		fmt.Println()
		fmt.Println("Press Enter to continue...")
		scanner.Scan()
		clearScreen()
	}

	for {
		fmt.Println("Regression Performance - Main Menu")
		fmt.Println("Please choose one of the following options:")
		fmt.Println("1. Calculate Performance Using Anscombe Quartet Set 1")
		fmt.Println("2. Calculate Performance Using Anscombe Quartet Set 2")
		fmt.Println("3. Calculate Performance Using Anscombe Quartet Set 3")
		fmt.Println("4. Calculate Performance Using Anscombe Quartet Set 4")
		fmt.Println("5. Calculate Performance Using ALL Anscombe Quartet Sets")
		fmt.Println("6. View Average Performance for All Tests in Current Session")
		fmt.Println("7. Graph Average Performance for All Tests in Current Session")
		fmt.Println("8. Help")
		fmt.Println("9. Configuration Menu")
		fmt.Println("0. Exit")

		var err error

		_, err = fmt.Scanf("%d\n", &choice)
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
			menuReset()
		case 2:
			set = "Two"
			experiment(set)
			fmt.Println()
			fmt.Println("Press Enter to continue...")
			scanner.Scan()
		case 3:
			set = "Three"
			experiment(set)
			fmt.Println()
			fmt.Println("Press Enter to continue...")
			scanner.Scan()
		case 4:
			set = "Four"
			experiment(set)
			fmt.Println()
			fmt.Println("Press Enter to continue...")
			scanner.Scan()
		case 5:
			fmt.Println("Running all Four Anscombe Quartets.")
			experiment("One")
			experiment("Two")
			experiment("Three")
			experiment("Four")
			fmt.Println()
			fmt.Println("Press Enter to continue...")
			scanner.Scan()
		case 6:
			fmt.Println("Calculating Performance on all runs:")
			calcPerformance()
			fmt.Println()
			fmt.Println("Press Enter to continue...")
			scanner.Scan()
		case 7:
			fmt.Println("Graphing Performance on all runs:")
			performanceGraph()
			fmt.Println()
			fmt.Println("Press Enter to continue...")
			scanner.Scan()
		case 8:
			help()
			fmt.Println()
			fmt.Println("Press Enter to continue...")
			scanner.Scan()
			clearScreen()
		case 9:
			optionsMenu()
		default:
			fmt.Println("Invalid choice! Please try again.")
		}

		if choice == 0 {
			return
		}
	}
}

func optionsMenu() {
	clearScreen()
	fmt.Println("Configuration Menu")
	fmt.Println()

	// determine which set to test on

	var choice int64 = -1

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Configuration Menu")
		fmt.Println("Please choose one of the following options:")
		fmt.Println("1. Set number of runs per test.")
		fmt.Println("2. Set number of runs per experiment.")
		fmt.Println("Press Enter to Return to Previous Menu")

		var err error

		_, err = fmt.Scanf("%d", &choice)
		if err != nil {
			choice = -1
		}

		switch choice {
		case 0:
			// Return to Previous Menu
			fmt.Println()
		case 1:
			fmt.Println("Runs per test determines number of times each model will be built. Times are then averaged and reported to application.")
			fmt.Println("Current number of runs per test:", nRuns)
			fmt.Println("Please enter new number and press enter.")
			_, err := fmt.Scanln(&nRuns)
			if err != nil {
				fmt.Println("Invalid input. Please try again.")
				continue
			}
			fmt.Println("Current number of runs per test:", nRuns)
			fmt.Println("Press Enter to continue...")
			scanner.Scan()
		case 2:
			fmt.Println("Runs per experiment determines number of times each experiment will be run in addition to number of times each test will be run (test x experiment = total number of runs)")
			fmt.Println("Current number of runs per experiment:", nExperiments)
			fmt.Println("Please enter new number and press enter.")
			_, err := fmt.Scanln(&nExperiments)
			if err != nil {
				fmt.Println("Invalid input. Please try again.")
				continue
			}
			fmt.Println("Current number of runs per experiment:", nExperiments)
			fmt.Println("Press Enter to continue...")
			scanner.Scan()
		default:
			return
		}

		if choice == 0 {
			return
		}
		//fmt.Println("Press Enter to continue...")
		//scanner.Scan() // Wait for user to press Enter

	}
}
