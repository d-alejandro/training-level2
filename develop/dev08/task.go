package main

import (
	"bufio"
	"d-alejandro/training-level2/develop/dev08/cmd"
	"fmt"
	"os"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать свой собственный UNIX-шелл-утилиту с
поддержкой ряда простейших команд:
- cd <args> - смена директории (в качестве аргумента могут
быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте
аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в
формате *такой-то формат*
Так же требуется поддерживать функционал fork/exec-команд
Дополнительно необходимо поддерживать конвейер на пайпах
(linux pipes, пример cmd1 | cmd2 | .... | cmdN).
*Шелл — это обычная консольная программа, которая будучи
запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись
ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный
сеанс поддерживается до тех пор, пока не будет введена
команда выхода (например \quit).
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	forkExecResultChannel := make(chan string, 1)
	handler := cmd.NewHandler(forkExecResultChannel)

	printFirstLine()

	for scanner.Scan() {
		command := scanner.Text()

		if command == "" {
			readChannelAndOutputResult(forkExecResultChannel)
			printFirstLine()
			continue
		}

		if command == "q" || command == "quit" || command == `\quit` || command == "exit" || command == `\exit` {
			break
		}

		response, err := handler.Execute(command)

		if err == nil {
			fmt.Println(response)
		} else {
			fmt.Println(err.Error())
		}

		printFirstLine()
	}

	close(forkExecResultChannel)
}

func printFirstLine() {
	fmt.Print("$ ")
}

func readChannelAndOutputResult(forkExecResultChannel <-chan string) {
	select {
	case result, ok := <-forkExecResultChannel:
		if ok {
			fmt.Println(result)
		}
	default:
	}
}
