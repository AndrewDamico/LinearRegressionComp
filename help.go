package main

import (
	"fmt"
	"os"
	"os/exec"
)

func help() {
	clearScreen()
	cmd := exec.Command("nano", "README.md")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error Python:", err)
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error Python:", err)
	}
	//fmt.Println(output)
}
