package main

import (
	"fmt"
	"net"
	"os"
	"unsafe"
)

func main() {
	const NetworkProtocolTCP = "tcp"
	const NetworkAddress = ":8080"

	listener, listeningError := net.Listen(NetworkProtocolTCP, NetworkAddress)
	if listeningError != nil {
		printErrorAndExit(listeningError)
	}
	defer closeListener(listener)

	for {
		connection, connectionError := listener.Accept()

		if connectionError != nil {
			fmt.Println(connectionError)
			closeConnection(connection)
			continue
		}

		go handleClient(connection)
	}
}

func handleClient(connection net.Conn) {
	defer closeConnection(connection)

	var inputBytes []byte

	for {
		readLength, readError := connection.Read(inputBytes)
		if readLength == 0 || readError != nil {
			fmt.Println(readError)
			break
		}

		inputString := string(inputBytes)

		outputString := fmt.Sprintf(
			"read length: %d, length: %d, size: %d",
			readLength,
			len(inputString),
			unsafe.Sizeof(inputString),
		)

		_, writeError := connection.Write([]byte(outputString))
		if writeError != nil {
			fmt.Println(writeError)
			break
		}
	}
}

func closeListener(listener net.Listener) {
	if err := listener.Close(); err != nil {
		printErrorAndExit(err)
	}
}

func closeConnection(connection net.Conn) {
	if err := connection.Close(); err != nil {
		fmt.Println(err)
	}
}

func printErrorAndExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}
