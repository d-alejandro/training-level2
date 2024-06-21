package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os"
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
	if strings.Contains(commandRow, "&") {
		commandStrings := receiver.splitBySeparatorAndTrimString(commandRow, "&")

		if len(commandStrings) != 2 || commandStrings[1] != "" {
			return "", errors.New("invalid fork/exec-commandRow")
		}

		commandSlice := receiver.splitByRegExprCommandRowAndTrimEndDelimiters(commandStrings[0])

		return receiver.runForkExecCommand(commandSlice)
	} else if strings.Contains(commandRow, "|") {
		commandStrings := receiver.splitBySeparatorAndTrimString(commandRow, "|")

		if receiver.isContainValues(commandStrings) {
			return "", errors.New("invalid pipes")
		}

		return receiver.executePipeCommand(commandStrings)
	} else if _, isContain := strings.CutPrefix(commandRow, "cd"); isContain {
		commandStrings := receiver.splitByRegExprCommandRowAndTrimEndDelimiters(commandRow)

		if len(commandStrings) != 2 {
			return "", errors.New("invalid cd command")
		}

		if err := os.Chdir(commandStrings[1]); err != nil {
			return "", err
		}

		return "Ok", nil
	}

	commandStrings := receiver.splitByRegExprCommandRowAndTrimEndDelimiters(commandRow)

	return receiver.executeCommand(commandStrings)
}

func (receiver *Handler) splitBySeparatorAndTrimString(command string, separator string) []string {
	commandStrings := strings.Split(command, separator)

	for index, commandString := range commandStrings {
		commandStrings[index] = strings.TrimSpace(commandString)
	}

	return commandStrings
}

func (receiver *Handler) runForkExecCommand(commandStrings []string) (string, error) {
	path, lookPathErr := exec.LookPath(commandStrings[0])

	if lookPathErr != nil {
		return "", lookPathErr
	}

	pid, forkExecErr := syscall.ForkExec(path, commandStrings, nil)

	if forkExecErr != nil {
		return "", forkExecErr
	}

	go receiver.waitForkExecCommandAndSendResult(pid)

	return strconv.Itoa(pid), nil
}

func (receiver *Handler) waitForkExecCommandAndSendResult(pid int) {
	var waitStatus syscall.WaitStatus = 0

	_, waitErr := syscall.Wait4(pid, &waitStatus, 0, nil)

	resultStatus := "Done"

	if waitErr != nil || waitStatus != 0 {
		resultStatus = "Error"
	}

	result := fmt.Sprintf("%d %s", pid, resultStatus)

	receiver.forkExecResultChannel <- result
}

func (receiver *Handler) splitByRegExprCommandRowAndTrimEndDelimiters(commandRow string) []string {
	regExpr := regexp.MustCompile(`('.*')|(".*")|(\S+)`)
	output := regExpr.FindAllString(commandRow, -1)

	for index, commandLine := range output {
		output[index] = strings.Trim(commandLine, `"'`)
	}

	return output
}

func (receiver *Handler) executeCommand(commandStrings []string) (string, error) {
	cmd, creationErr := receiver.createCommand(commandStrings)

	if creationErr != nil {
		return "", creationErr
	}

	result, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	return string(result), nil
}

func (receiver *Handler) createCommand(commandStrings []string) (*exec.Cmd, error) {
	path, lookPathErr := exec.LookPath(commandStrings[0])

	if lookPathErr != nil {
		return nil, lookPathErr
	}

	return exec.Command(path, commandStrings[1:]...), nil
}

func (receiver *Handler) isContainValues(commandStrings []string) bool {
	return commandStrings[0] == "" || commandStrings[len(commandStrings)-1] == ""
}

func (receiver *Handler) executePipeCommand(commandStrings []string) (string, error) {
	var outputBuffer bytes.Buffer

	commandStringsLength := len(commandStrings)
	commandSlice := make([]*exec.Cmd, commandStringsLength)

	for index, command := range commandStrings {
		commandArgs := receiver.splitByRegExprCommandRowAndTrimEndDelimiters(command)
		cmd, cmdError := receiver.createCommand(commandArgs)

		if cmdError != nil {
			return "", cmdError
		}

		commandSlice[index] = cmd

		if index > 0 {
			commandSlice[index].Stdin, _ = commandSlice[index-1].StdoutPipe()
		}
	}

	commandSlice[commandStringsLength-1].Stdout = &outputBuffer
	commandSlice[commandStringsLength-1].Stderr = &outputBuffer

	for index := commandStringsLength - 1; index > 0; index-- {
		if startError := commandSlice[index].Start(); startError != nil {
			return "", startError
		}
	}

	if runError := commandSlice[0].Run(); runError != nil {
		return "", runError
	}

	for index := 1; index < commandStringsLength; index++ {
		if waitError := commandSlice[index].Wait(); waitError != nil {
			return "", waitError
		}
	}

	return string(outputBuffer.Bytes()), nil
}
