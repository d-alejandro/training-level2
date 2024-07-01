package main

import (
	"d-alejandro/training-level2/develop/dev10/socket"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	const NetworkProtocolTCP = "tcp"
	const NetworkAddress = "localhost:8080"

	server, serverError := socket.NewServer(NetworkProtocolTCP, NetworkAddress)

	if serverError != nil {
		printErrorAndExit(serverError)
	}

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	<-channel

	if err := server.Stop(); err != nil {
		printErrorAndExit(err)
	}
}

func printErrorAndExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}
