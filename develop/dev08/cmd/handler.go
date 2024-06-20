package cmd

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (receiver *Handler) Execute(command string) (string, error) {
	switch {
	case strings.Contains(command, "&"):
		commandStrings := strings.Split(command, "&")

		for index, commandString := range commandStrings {
			commandStrings[index] = strings.TrimSpace(commandString)
		}

		if len(commandStrings) != 2 || commandStrings[1] != "" {
			return "", errors.New("invalid fork/exec-command")
		}

		return receiver.runForkExecCommand(commandStrings[0])

	case strings.Contains(command, "|"):
	}

	return command, nil
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

	return strconv.Itoa(pid), nil
}
