package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
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

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	const NetworkProtocolTCP = "tcp"
	address := net.JoinHostPort(arguments[0], arguments[1])
	connection, connectionError := net.DialTimeout(NetworkProtocolTCP, address, *timeoutDurationFlag)

	if connectionError != nil {
		printMessageAndExit(connectionError.Error())
	}
	defer func() {
		if err := connection.Close(); err != nil {
			printMessageAndExit(err.Error())
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		_ = text
		//if text == "" {
		//	continue
		//}
		//
		//if text == "q" || text == "quit" || text == "exit" {
		//	break
		//}
	}
	waitGroup.Done()
	waitGroup.Wait()
}

func printMessageAndExit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
