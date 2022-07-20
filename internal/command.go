package internal

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

// ClearTerminal - print out statement that will force a terminal to clear
func ClearTerminal() {
	fmt.Print("\033[H\033[2J") // taken from: https://stackoverflow.com/a/22892171
}

// ExecuteCommand - execute a command with args
func ExecuteCommand(logger *log.Logger, command string, commandArgs []string) ([]byte, error) {
	// echo simulated input
	fmt.Println("> ", command, strings.Join(commandArgs, " "))

	// ignore error as it gets captured
	output, _ := exec.Command(command, commandArgs...).CombinedOutput() // always allow the command to possibly error

	// echo stringified output (apparently autocolorizes)
	fmt.Println(string(output))

	// wait for element to render before proceeding
	logger.Println("Finished executing command", command, commandArgs)
	time.Sleep(350 * time.Millisecond) // waiting because the parent terminal process may not have finished rendering

	return output, nil
}
