package main

import (
	"bufio"
	"d-alejandro/training-level2/develop/dev10/socket"
	"flag"
	"fmt"
	"os"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные из сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться.
Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему серверу, программа должна завершаться через timeout.
*/

func main() {
	const TimeoutDuration = 10 * time.Second
	timeoutDurationFlag := flag.Duration("timeout", TimeoutDuration, "timeout duration")

	flag.Parse()

	if *timeoutDurationFlag < 1*time.Second {
		printMessageAndExit("flag value is less than one.")
	}

	arguments := flag.Args()

	if len(arguments) != 2 {
		printMessageAndExit("wrong number of parameters.")
	}

	connectionController, connectionControllerError := socket.NewConnectionController(
		arguments[0],
		arguments[1],
		*timeoutDurationFlag,
	)
	if connectionControllerError != nil {
		printMessageAndExit(connectionControllerError.Error())
	}

	client, clientError := socket.NewClient(arguments[0], arguments[1], *timeoutDurationFlag)
	if clientError != nil {
		printMessageAndExit(clientError.Error())
	}

	go func() {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			text := scanner.Text()

			response, err := client.Send(text)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println(response)
		}

		connectionController.Stop()
	}()

	<-connectionController.GetQuitChannel()

	client.Stop()

	connectionController.Stop()
}

func printMessageAndExit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
