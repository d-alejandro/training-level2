package cmd

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

type Handler struct {
	forkExecResultChannel chan<- string
}

func NewHandler(forkExecResultChannel chan<- string) *Handler {
	return &Handler{forkExecResultChannel}
}

func (receiver *Handler) Execute(command string) (string, error) {
	switch {
	case strings.Contains(command, "&"):
		commandStrings := receiver.splitAndTrimStrings(command)

		if len(commandStrings) != 2 || commandStrings[1] != "" {
			return "", errors.New("invalid fork/exec-command")
		}

		return receiver.runForkExecCommand(commandStrings[0])

	case strings.Contains(command, "|"):
	}

	return command, nil
}

func (receiver *Handler) splitAndTrimStrings(command string) []string {
	commandStrings := strings.Split(command, "&")
	for index, commandString := range commandStrings {
		commandStrings[index] = strings.TrimSpace(commandString)
	}
	return commandStrings
}

func (receiver *Handler) runForkExecCommand(command string) (string, error) {
	commandSlice := strings.Split(command, " ")

	path, lookPathErr := exec.LookPath(commandSlice[0])
	if lookPathErr != nil {
		return "", lookPathErr
	}

	pid, forkExecErr := syscall.ForkExec(path, commandSlice, nil)

	if forkExecErr != nil {
		return "", forkExecErr
	}

	go func() {
		var waitStatus syscall.WaitStatus = 0

		_, waitErr := syscall.Wait4(pid, &waitStatus, 0, nil)

		resultStatus := "Done"

		if waitErr != nil || waitStatus != 0 {
			resultStatus = "Error"
		}

		result := fmt.Sprintf("%d %s", pid, resultStatus)

		receiver.forkExecResultChannel <- result
	}()

	return strconv.Itoa(pid), nil
}
