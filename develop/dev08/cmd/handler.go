package cmd

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
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

func (receiver *Handler) Execute(commandRow string) (string, error) {
	switch {
	case strings.Contains(commandRow, "&"):
		commandStrings := receiver.splitAndTrimString(commandRow)

		if len(commandStrings) != 2 || commandStrings[1] != "" {
			return "", errors.New("invalid fork/exec-commandRow")
		}

		return receiver.runForkExecCommand(commandStrings[0])
	case strings.Contains(commandRow, "|"):
	}

	commandSlice := receiver.splitCommandRow(commandRow)

	path, lookPathErr := exec.LookPath(commandSlice[0])
	if lookPathErr != nil {
		return "", lookPathErr
	}

	cmd := exec.Command(path, commandSlice[1:]...)

	result, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	return string(result), nil
}

func (receiver *Handler) splitAndTrimString(command string) []string {
	commandStrings := strings.Split(command, "&")
	for index, commandString := range commandStrings {
		commandStrings[index] = strings.TrimSpace(commandString)
	}
	return commandStrings
}

func (receiver *Handler) runForkExecCommand(command string) (string, error) {
	commandSlice := receiver.splitCommandRow(command)

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

func (receiver *Handler) splitCommandRow(commandRow string) []string {
	regExpr := regexp.MustCompile(`('.*')|(".*")|(\S+)`)
	return regExpr.FindAllString(commandRow, -1)
}
